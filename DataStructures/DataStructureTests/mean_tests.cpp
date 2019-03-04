#include "stdafx.h"
#include "CppUnitTest.h"
#include "../MeanCalculatingStructure/mean_structure.cpp"
using namespace Microsoft::VisualStudio::CppUnitTestFramework;


TEST_CLASS(mean_tests)
{
public:

	TEST_METHOD(TestMethod1)
	{
		mean_structure* m_s = new mean_structure();
		double p = 5.0;
		m_s->insert(5.6);
		m_s->insert(5.4);
		
		double median = m_s->get_median();
		delete m_s;
		Assert::AreEqual(5.5, median);
		//m_s->insert(p);
		//m_structure.get_median();
		// TODO: Your test code here
	}

};
