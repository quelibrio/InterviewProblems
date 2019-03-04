#include "CppUnitTest.h"
#include "../MeanCalculatingStructure/mean_structure.cpp"
using namespace Microsoft::VisualStudio::CppUnitTestFramework;


TEST_CLASS(mean_tests)
{
	mean_structure* m_s;
public:
	TEST_METHOD_INITIALIZE(initialize)
	{
		m_s = new mean_structure();
	}

	TEST_METHOD_CLEANUP(clenUp)
	{
		m_s->clear();
	}

	TEST_METHOD(testMedianOneBigNumber)
	{
		m_s->insert(2 ^ 30);
		double median = m_s->get_median();
		Assert::AreEqual((2 ^ 30)*1.0, median);
	}

	TEST_METHOD(testNoNumbers)
	{
		double median = m_s->get_median();
		Assert::AreEqual((2 ^ 31)*(-1.0), median);
	}

	TEST_METHOD(testMedianIntegers)
	{
		m_s->insert(6);
		m_s->insert(8);
		double median = m_s->get_median();
		
		Assert::AreEqual(7.0, median);
	}
	TEST_METHOD(testMedianTwoNumber)
	{
		m_s->insert(5.6);
		m_s->insert(5.4);
		double median = m_s->get_median();
		Assert::AreEqual(5.5, median);
	}

	TEST_METHOD(testMedianFiveNumbers)
	{
		mean_structure* m_s = new mean_structure();
		m_s->insert(75);
		m_s->insert(82);
		m_s->insert(89);
		m_s->insert(91);
		m_s->insert(102);
		double median = m_s->get_median();
		Assert::AreEqual(89.0, median);
	}

	TEST_METHOD(testMedianNineNumbers)
	{
		m_s->insert(20.5);
		m_s->insert(26.1);
		m_s->insert(32.5);
		m_s->insert(18.9);
		m_s->insert(33.4);
		m_s->insert(29.7);
		m_s->insert(19.8);
		m_s->insert(25.6);
		m_s->insert(34.3);
		double median = m_s->get_median();
		Assert::AreEqual(26.1, median);
	}

	TEST_METHOD(testMedianThousandsOfNumbersReverse)
	{
		for (double i = 99999.0; i > 0; i--)
		{
			m_s->insert(i);
		}
		double median = m_s->get_median();
		Assert::AreEqual(50000.0, median);
	}

	TEST_METHOD(testMedianThousandsOfNumbersMultipleInsert)
	{
		for (double i = 999.0; i > 0; i--)
		{
			m_s->insert(i);
		}
		double median = m_s->get_median();
		Assert::AreEqual(500.0, median);

	    median = m_s->get_median();
		for (double i = 999.0; i > 0; i--)
		{
			m_s->insert(i);
		}
		Assert::AreEqual(500.0, median);

		m_s->insert(-1.0);
		m_s->insert(-1.0);
		m_s->insert(-1.0);
	    median = m_s->get_median();
		Assert::AreEqual(499.0, median);
	}
};
