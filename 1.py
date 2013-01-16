#!/usr/bin/python

import math
import operator

def lower_floor_div(x,y):
    floor = x/y
    if (x % y) == 0:
        floor = floor - 1
    return floor

def sum_multiples_less_than(factor, y):
    greatest_factor = lower_floor_div(y, factor)
    return (factor * ((greatest_factor + 1) * greatest_factor)/2)

def soln(factors, top):
    sums = [sum_multiples_less_than(factor, top) for factor in factors]
    return reduce(operator.add, sums)
    
    
