#include "stdafx.h"
#include <queue>
#include <functional>
#include <iostream>
#include <cmath>   

using namespace std;
// Creates a min heap 
priority_queue <double> max_pq;
priority_queue <double, vector<double>, greater<double>> min_pq;

bool isEmpty()
{
	return max_pq.size() == 0 && min_pq.size() == 0;
}

double getMedian() {
	//if total size is even, then median is the 2 middle elements' average
	if (max_pq.size() == min_pq.size())
	{
		return (max_pq.top() + min_pq.top()) / 2;
	}
	//otherwise median is middle element then the root of the heap with one element more is the median
	else if (max_pq.size() > min_pq.size())
	{
		return (double)max_pq.top();
	}
	else
	{
		return (double)min_pq.top();
	}
}

void balanceTreeSize()
{
	//if heaps' sizes differ by 2, then we need to redistribute elements
	//from the bigger heap to the smaller
	if (std::abs((int)(max_pq.size() - min_pq.size())) > 1)
	{
		if (max_pq.size() > min_pq.size())
		{
			min_pq.push(max_pq.top());
			max_pq.pop();
		}
		else {
			max_pq.push(min_pq.top());
			min_pq.pop();
		}
	}
}

void insert(int n) {
	// push to the minimal priority queue if we don't have elements
	if (isEmpty())
	{
		min_pq.push(n);
	}
	else
	{
		// if n is less than or equal to current median, add to maxheap
		if (n <= getMedian())
		{
			max_pq.push(n);
		}
		// if n is greater than current median, add to min heap
		else
		{
			min_pq.push(n);
		}
	}
	//fix balance of priority queue size after insertion
	balanceTreeSize();
}

int main()
{
	insert(82);
	insert(102);
	insert(75);
	insert(91);
	insert(89);
	insert(91);
	cout << getMedian();

	return 0;
}
