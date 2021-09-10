using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;

namespace problem_6_sum_square_difference
{
    class Program
    {
        static void Main(string[] args)
        {
            Calculator calculator = new Calculator(100);

            var sumOfSquares = calculator.GetSumOfSquares();
            var squareOfSums = calculator.GetSquareOfSums();
            
            var difference = Calculator.GetDifferenceBetweenNumbers(squareOfSums, sumOfSquares);
            
            Console.WriteLine("DIFFERENCE :");
            Console.WriteLine(difference);
        }
    }

    class Calculator
    {
        private readonly IEnumerable<int> _numbers;
        public Calculator(int sizeOfDataset)
        {
            this._numbers = Calculator.GetNaturalNumbers(sizeOfDataset);
        }

        private static IEnumerable<int> GetNaturalNumbers(int sizeOfList)
        {
            var listOfNaturalNumbers = new List<int>();

            for (var i = 1; i <= sizeOfList; i++)
            {
                listOfNaturalNumbers.Add(i);
            }

            return listOfNaturalNumbers;
        }
        
        public static double GetDifferenceBetweenNumbers(double argumentOne, double argumentTwo)
        {
            return argumentOne - argumentTwo;
        }

        public double GetSquareOfSums()
        {
            var sum = this._numbers.Sum();
            return Math.Pow(sum, 2);
        }

        public double GetSumOfSquares()
        {
            return this._numbers.Sum(number => Math.Pow(number, 2));
        }
    }
}
