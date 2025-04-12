class Solution:
    def teleportationPortals(self, NP1P2, A1AN):
        n, p1, p2 = NP1P2[0], min(NP1P2[1], NP1P2[2]), max(NP1P2[1], NP1P2[2])
        tp_lvls_before = A1AN
        tp_lvls = A1AN[:]

        if p1 not in tp_lvls:
            tp_lvls.append(p1)
        
        if p2 not in tp_lvls:
            tp_lvls.append(p2)
    
        tp_lvls_sorted = sorted(tp_lvls)
        p1_idx = tp_lvls_sorted.index(p1)
        p2_idx = tp_lvls_sorted.index(p2)

        max_lvl = int(tp_lvls_sorted[len(tp_lvls_sorted)-1])+1
        dp_left_p1 = [0] * max_lvl
        dp_right_p1 = [0] * max_lvl
        dp_left_p2 = [0] * max_lvl
        dp_right_p2 = [0] * max_lvl
        
        i=0
        if (p1_idx > 0):
            i = p1_idx-1
            max_v = 0
            while (i>=0):
                max_v = max(max_v, abs(tp_lvls_sorted[i]-tp_lvls_sorted[i+1]))
                dp_left_p1[tp_lvls_sorted[i]] = max_v
                i-=1

        i = p1_idx+1
        max_v = 0
        while (tp_lvls_sorted[i]<p2):
            max_v = max(max_v, abs(tp_lvls_sorted[i-1]-tp_lvls_sorted[i]))
            dp_right_p1[tp_lvls_sorted[i]] = max_v
            i+=1

        if (p2_idx > 0):
            i = p2_idx-1
            max_v = 0
            while (i>p1):
                max_v = max(max_v, abs(tp_lvls_sorted[i]-tp_lvls_sorted[i+1]))
                dp_left_p2[tp_lvls_sorted[i]] = max_v
                i-=1

        i = p2_idx+1
        max_v = 0
        while (i<len(tp_lvls)):
            max_v = max(max_v, abs(tp_lvls_sorted[i-1]-tp_lvls_sorted[i]))
            dp_right_p2[tp_lvls_sorted[i]] = max_v
            i+=1
        
        min_suit_lvls = [0] * n
        for i in range(n):
            if tp_lvls_before[i] < p1:
                min_suit_lvls[i] = dp_left_p1[tp_lvls_before[i]]
            elif tp_lvls_before[i] > p2:
                min_suit_lvls[i] = dp_right_p2[tp_lvls_before[i]]
            else:
                if all((dp_right_p1[tp_lvls_before[i]], dp_left_p2[tp_lvls_before[i]])):
                    min_suit_lvls[i] = min(dp_right_p1[tp_lvls_before[i]], dp_left_p2[tp_lvls_before[i]])
                else:
                    if not dp_right_p1[tp_lvls_before[i]]:
                        min_suit_lvls[i] = dp_left_p2[tp_lvls_before[i]]
                    else:
                        min_suit_lvls[i] = dp_right_p1[tp_lvls_before[i]]
        
        return " ".join(map(str, min_suit_lvls))

sol = Solution()
arg1 = list(map(int, [1, 1110, 1000000000]))
arg2 = list(map(int, [1000000000]))
result = sol.teleportationPortals(arg1,arg2)
print(result)