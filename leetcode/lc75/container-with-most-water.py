class Solution:
    from typing import List
    def maxArea(self, height: List[int]) -> int:
        max_area = 0
        left = 0
        right = len(height)-1
        max_left = height[left]
        max_right = height[right]

        while (left < right):
            print("{} {}\n".format(left, right))
            length = abs(left-right)
            _height = min(height[left], height[right])
            
            area = length*_height
            max_area = max(area, max_area)

            if height[left] < height[right]:
                left+=1
            else:
                right-=1

        return max_area

sol = Solution()
result = sol.maxArea([1,3,2,5,25,24,5])
print(result)