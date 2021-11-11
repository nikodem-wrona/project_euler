from math import *

DEFAULT_EXPONENT = 2;

def getPowOfNumber(base):
  return pow(base, DEFAULT_EXPONENT)

def checkPythagoreanTriplet(numberOne, numberTwo, numberThree):
  sum_left = getPowOfNumber(numberOne) + getPowOfNumber(numberTwo)
  
  if (sum_left == getPowOfNumber(numberThree)):
    return True
  else:
    return False


def getResult():
  results = []

  for a in xrange(1000):
    for b in xrange(1000):
      c = 1000 - ( a + b )
      condition = a > 0 and b > 0 and c > 0

      if c > 0 and (a + b + c) == 1000 and condition:
        is_pythagorean_triplet = checkPythagoreanTriplet(a, b, c)
        if (is_pythagorean_triplet):
          results.append(a * b * c)
  
  return results

result = getResult()
print("Result : ", result)