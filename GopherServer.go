//Gopher translator service

//Gophers are friendly creatures but it’s not that easy to communicate with them. They have their own language and they don’t understand English.
//Create a program that starts http server. This server should be able to translate English words into words in the gophers' language. Don't worry, the gophers' language is pretty easy.
//The language that the gophers speak is a modified version of English and has a few simple rules.
//1.	If a word starts with a vowel letter, add prefix “g” to the word (ex. apple => gapple)
//2.	If a word starts with the consonant letters “xr”, add the prefix “ge” to the begging of the word. Such words as “xray” actually sound in the beginning with vowel sound as you pronounce them so a true gopher would say “gexray”.
//3.	If a word starts with a consonant sound, move it to the end of the word and then add “ogo” suffix to the word. Consonant sounds can be made up of multiple consonants, a.k.a. a consonant cluster (e.g. "chair" -> "airchogo”).
//4.	If a word starts with a consonant sound followed by "qu", move it to the end of the word, and then add "ogo" suffix to the word (e.g. "square" -> "aresquogo").
//Your program should accept one command line argument “—port” which is the port that the server is running.
//Your http server should have the following endpoints:
//5.	POST “/word” - by given English word, the server should return the word’s translation in gopher language. It should accept json data in the format {“english-word”:”<a single English word>”} and should return json data in the format {“gopher-word”:”<translated version of the given word>”}
//6.	(OPTIONAL) POST “/sentence” - by given English sentence (in which each whitespace separated sequence counts as single word) the server should return the sentence translation in gopher language.  It should accept json data in the format {“english-sentence”:”<sentence of English words>”} and return {“gopher-sentence”:”<translated version of the given sentence>”} Assume that every sentence ends with dot, question or exclamation mark.
//7.	(OPTIONAL) GET “/history” - should return each English word or sentence that was given to the server from the time the server was started along with its translation in gopher language. The output should look like {“history”:[{“apple”:”gapple”},{“my”:”ymogo”},….]}  The returned array should be ordered alphabetically ascending by the English word/sentence.
//Please don’t confuse the gophers as they don’t understand shortened versions of words or apostrophes. So don’t use words like - “don’t”, “shouldn’t”, etc. Even translated they still won’t understand you so skip them in your solution. 
//It is necessary your code to compile. 


/// Usage 'go run gopher_server.go 8081'
/// History implemented with map
/// Insertion by a word or sentence - O(1)
/// Query - O(NlogN) where N is the number of words/sentences inserted
/// Posible optimization with an ordered heap: Insert O(logN) Query(logN)

//PS. This solution recieved feedback from FT that is in 1 file and got rejected. Dont use directly.


package main

import (
    "encoding/json"
    "log"
    "net/http"
	"strings"
    "os"
    "sort"
    "sync"
    //"fmt"
    //"io/ioutil"
    "testing"
)

//Structs for organizing HTTP request bodies
type WordStruct struct {
    Word string `json:"english-word"`
}

type WordStructOutput struct{
    Word string `json:"gopher-word"`
}

type SentenceStruct struct {
    Word string `json:"english-sentence"`
}

type SentenceStructOutput struct{
    Word string `json:"gopher-sentence"`
}


//In memory hitory of the translation map
var translationMap map[string]string

//Mutext to prevent concurrrent writes to the translation map
var mutex = &sync.Mutex{}

//HTTP Post /word
func wordHandler(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var t WordStruct
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
        return
    }
    word:=string(t.Word)
    translatedWord:=translateWord(word)

	rw.Header().Set("Content-Type", "application/json") 
	words := WordStructOutput {
            Word: translatedWord,
        }
    updateHistory(word, translatedWord)
    json.NewEncoder(rw).Encode(words) 
}

//HTTP Post /sentence
func sentenceHandler(rw http.ResponseWriter, req *http.Request) {
    decoder := json.NewDecoder(req.Body)
    var t SentenceStruct
    err := decoder.Decode(&t)
    if err != nil {
        return
        panic(err)
    }

    word:=string(t.Word)
    translatedWord:=translateSentence(word)

	rw.Header().Set("Content-Type", "application/json") 
	words := SentenceStructOutput {
            Word: translatedWord,
        } 
    updateHistory(word, translatedWord)
    json.NewEncoder(rw).Encode(words) 
}

//HTTP Get /history
func historyHandler(rw http.ResponseWriter, req *http.Request){
    keys := make([]string, len(translationMap))
    i := 0
    for k := range translationMap {
        keys[i] = k
        i++
    }
    sort.Strings(keys)
    jsonString, _ := json.Marshal(translationMap)

    var historyOutput map[string]string
    historyOutput = make(map[string]string)
    historyOutput["history"]=string(jsonString)
    historyWrapped, _ := json.Marshal(historyOutput)
    json.NewEncoder(rw).Encode((string(historyWrapped))) 
}

//Translate a single word into the gopher language
func translateWord(word string) string {
    //Predefined are vowels and consonants to implement language rules
	vowels := []string{"a", "o", "u", "e", "i", "y"}
    consonantSounds := []string{"tch","dge", "ch", "th", "es", "sh", 
        "ge", "ll", "ng", "ry", "p", "b", "t", 
        "d", "c", "g", "f", "v", "s", "z", "s", 
        "h", "m", "n", "l", "r", "w", "y", "q"}

    translatedWord:=""
	first_letter:=string(word[0])

    //Words starts with a vowel letter
	if stringInSlice(first_letter, vowels) {
		translatedWord="g"+word
        return translatedWord
	} else if strings.HasPrefix(word, "xr") {
        //Word starts with a xr
		translatedWord="ge"+word
        return translatedWord
	} else {
        for _, consonantSound := range consonantSounds {
            if strings.HasPrefix(word, consonantSound){
            translatedWord=processConsonantBegining(word, consonantSound)
            return translatedWord
            }
        }
    }
    return translatedWord
}

//Word starts with a consonant sound
func processConsonantBegining(word string, consonantSound string) string {  
    translatedWord:=""
    prefixLength:=len(consonantSound)
    //Check if we have enough letters to match the first rule
    if len(word)>prefixLength+2{
        nextTwoLetters:=word[prefixLength:prefixLength+2]
        //Starts with qu
        if nextTwoLetters=="qu"  {
            firstThreeLetters:=word[0:prefixLength+2]
            remainingWord:=word[prefixLength+2:]
            translatedWord=remainingWord+firstThreeLetters+"ogo"
        }else{
            //Just starts with a consonant sound
            remainingWord:=word[prefixLength:]
            translatedWord=remainingWord+consonantSound+"ogo"
            
        }
    }else{
        //Just starts with a consonant sound
        remainingWord:=word[prefixLength:]
        translatedWord=remainingWord+consonantSound+"ogo"
    }
    return translatedWord
}

//Translate a sentence of multiple words
func translateSentence(sentence string) string{
    final_translation:=""
    sentence=cleanPunctuationAndCase(sentence)
	words := strings.Fields(sentence)
    for _, word := range words {
        translated_word:=translateWord(word)
        final_translation=final_translation+" "+translated_word
    }
    return final_translation[1:]
}

//Clearn punctuation. Lowercase captial letters
func cleanPunctuationAndCase(sentence string) string{
    sentence = strings.ReplaceAll(sentence, ",", " ")
    sentence = strings.ReplaceAll(sentence, ".", " ")
    sentence = strings.ReplaceAll(sentence, ";", " ")
    sentence = strings.ReplaceAll(sentence, "!", " ")
    sentence = strings.ReplaceAll(sentence, "?", " ")
    return strings.ToLower(sentence)
}

//Lock and unlock to write in the map that could be used in many paralel requests
func updateHistory(to_translate string, translation string){
    mutex.Lock()
    translationMap[to_translate]=translation
    mutex.Unlock()
}

//Is string element in slices
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

///Program starting point
//First argument is expected to be the port number
func main() {
    argsWithoutProg := os.Args[1]
    
    //Initialize translation map used for hitory lookups
    translationMap = make(map[string]string)

    http.HandleFunc("/word", wordHandler)
    http.HandleFunc("/sentence", sentenceHandler)
    http.HandleFunc("/history", historyHandler)
    
    log.Fatal(http.ListenAndServe(":"+argsWithoutProg, nil))
}

//Additional testing methods which are not used by the server
func TestTranslationSentence1(t *testing.T) {
    sentence:="apple xray chair square"
    translation:="gapple gexray airchogo aresquogo"
    translatedSenternce:=translateSentence(sentence)
    if translatedSenternce != translation {
        t.Errorf("Incorrect translation. %s %s %s", sentence, translation, translateSentence)
    }
}

func TestTranslationSentence2(t *testing.T) {
    sentence:="code one juice all xray xro g xr A"
    translation:="odecogo gone  gall gexray gexro gogo gexr ga"
    translatedSenternce:=translateSentence(sentence)
    if translatedSenternce != translation {
        t.Errorf("Incorrect translation. %s %s %s", sentence, translation, translateSentence)
    }
}

func TestTranslationSentence3(t *testing.T) {
    sentence:="gopher yep go"
    translation:="ophergogo gyep ogogo"
    translatedSenternce:=translateSentence(sentence)
    if translatedSenternce != translation {
        t.Errorf("Incorrect translation. %s %s %s", sentence, translation, translateSentence)
    }
}

func TestTranslationSentence4(t *testing.T) {
    sentence:="apple,xray. chair!square"
    translation:="gapple gexray airchogo aresquogo"
    translatedSenternce:=translateSentence(sentence)
    if translatedSenternce != translation {
        t.Errorf("Incorrect translation. %s %s %s", sentence, translation, translateSentence)
    }
}

func TestTranslationSentence5(t *testing.T) {
    sentence:="zzqall golf are VeRy Fast Cars"
    translation:="zqallzogo olfgogo gare eryvogo astfogo arscogo"
    translatedSenternce:=translateSentence(sentence)
    if translatedSenternce != translation {
        t.Errorf("Incorrect translation. %s %s %s", sentence, translation, translateSentence)
    }
}
