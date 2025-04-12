class Solution:
    from typing import List
    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        n = len(rooms)
        if (n <= 1):
            return True

        stack = [0]
        visited = [False] * n
        visited[0] = True

        while(stack):
            room_idx = stack.pop()
            for key in rooms[room_idx]:
                if not visited[key]:
                    stack.append(key)
                    visited[key] = True
                if all(visited):
                    return True
        
        return all(visited)

        
sol = Solution()
result = sol.canVisitAllRooms([[1,3],[3,0,1],[2],[0]])
print(result)