#include <queue>
#include <functional>
#include <iostream>
#include <cmath> 
#include "mean_structure.h"
using namespace std;

const double not_a_number = (-1.0)*(2 ^ 31);
// Creates a maximum order heap 
priority_queue <double> max_pq;
// creates a minimum order heap with the generic constructor
priority_queue <double, vector<double>, greater<double>> min_pq;

bool mean_structure::is_empty()
{
	return max_pq.size() == 0 && min_pq.size() == 0;
}

void mean_structure::clear()
{
	max_pq = priority_queue<double>();
	min_pq = priority_queue<double, vector<double>, greater<double>>();
}

double mean_structure::get_median()
{
	if(max_pq.size() == 0 && min_pq.size() == 0)
	{
		return not_a_number;
	}
	//if total size is even, then median is the 2 middle elements' average
	else if (max_pq.size() == min_pq.size())
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

void mean_structure::balance_tree_size()
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

void  mean_structure::insert(double n) {
	// push to the minimal priority queue if we don't have elements
	if (mean_structure::is_empty())
	{
		min_pq.push(n);
	}
	else
	{
		// if n is less than or equal to current median, add to maxheap
		if (n <= get_median())
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
	balance_tree_size();
}

int main()
{
	mean_structure* structure = new mean_structure();
	structure->insert(82);
	structure->insert(102);
	structure->insert(75);
	structure->insert(91);
	structure->insert(89);
	structure->insert(91);
	cout << structure->get_median();
	//First run -> median should be 90
	return 0;
}
