class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        answer = [1]*n
        prefix, postfix = 1, 1

        for i in range(n):
            answer[i] *= prefix
            prefix *= nums[i]
            answer[n-i-1] *= postfix
            postfix *= nums[n-i-1]
            
        return answer