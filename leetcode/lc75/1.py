#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'lampuDanTombolSpesial' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. INTEGER N
#  2. INTEGER_ARRAY L
#

def lampuDanTombolSpesial(N, L):
    return print("{} {}".format("YES", 5))

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    T = int(input().strip())

    for T_itr in range(T):
        N = int(input().strip())

        L = list(map(int, input().rstrip().split()))

        result = lampuDanTombolSpesial(N, L)

        fptr.write(result + '\n')

    fptr.close()
