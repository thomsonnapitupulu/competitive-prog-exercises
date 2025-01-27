class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        maxlen = 0
        word1len = len(word1)
        word2len = len(word2)
        maxlen = word1len
        merged = ''

        if word2len > maxlen:
            maxlen = len(word2)

        for i in range(maxlen):
            if i < word1len:
                merged += word1[i]
            if i < word2len:
                merged += word2[i]
        
        return merged