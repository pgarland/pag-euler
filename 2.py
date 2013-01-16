#!/usr/bin/python

import math

phi = (1 + math.sqrt(5))/2

def fib(n):
    return int(round((math.pow(phi, n+1) - math.pow(1 - phi, n+1))/ math.sqrt(5)))

def fibs():
    x = 0
    while True:
        yield fib(x)
        x = x+1

def soln():
    sum = 0
    for f in fibs():
        if f > 2000000:
            break
        if f % 2 == 0:
            sum += f
    return sum



