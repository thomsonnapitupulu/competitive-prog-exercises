class Solution:
    def cetekin(l):
        return (l+1)%2

    def lampuDanTombolSpesial(N, L):
        i = 0
        times=0
        result = ""
        
        while(i<len(L)-2):
            if (i==len(L)-3):
                if (all([L[i], L[i+1], L[i+2]])):
                    result = "YES {}".format(times)
                    return result
                elif (all([cetekin(L[i]), cetekin(L[i+1]), cetekin(L[i+2])])):
                    result = "YES {}".format(times+1)
                    return result
                else:
                    result = "NO -1"
                    return result

            if (L[i] == 0):
                L[i], L[i+1], L[i+2] = cetekin(L[i]), cetekin(L[i+1]), cetekin(L[i+2])
                times+=1
        
sol = Solution()
lamps = map(int, [1, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0, 0, 1])
result = sol.lampuDanTombolSpesial(13,lamps)
print(result)