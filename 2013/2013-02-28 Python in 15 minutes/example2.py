# -*- coding: utf-8 -*-
# numbers
print 1
print 3.2
print 5j+7  # complex  
print 6e-5   # 6 * 10^-5 
import math
print math.sqrt

def static(f):
	def g(self, *a):
		return f(*a)
	return g


class f:
	z = 52
	def _api(s):
		pass # do sth here
	
	@static
	def __007(n):
		return "I am Bond, James Bond {}".format(n)
		
	def __repr__(self):
		return "I am F {} {}".format(self.z, self.__007(7))

print repr(f())
	
print (3**2 + 4**2)**.5   # this is 5

# booleans
print True
print False

# strings
print "I am a beautiful string"
print u"我也是！"

# lists
print [1, 2, 4]
print ["can have", 4, "mixed types", ["even another list"]]

# tuples
print ("similar do list", "but immutable")

# dictionary (hash map)
print {"key1": "value", "another key": 52}

# sets
print set(["each value once", "each value once"])

# list comprehension
print [2*x for x in range(5)]

# dict comprehension
print {"key{}".format(k): k**2 for k in range(6)}
