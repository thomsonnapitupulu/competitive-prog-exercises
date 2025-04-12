class Solution:
    from collections import deque
    def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
        q = deque([(entrance[0], entrance[1], 0)])
        max_col, max_row = len(maze)-1, len(maze[0])-1 # define the boundaries
        directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
        maze[entrance[0]][entrance[1]] = '+' # mark entrance as visited

        while (q):
            c, r, nsteps = q.popleft()
            for next_move in directions: 
                new_c, new_r = c + next_move[0], r + next_move[1]
                if 0 <= new_c <= max_col and 0 <= new_r <= max_row and maze[new_c][new_r] == '.':
                    if new_c == 0 or new_c == max_col or new_r == 0 or new_r == max_row:
                        return nsteps+1
                    maze[new_c][new_r] = '+'
                    q.append((new_c, new_r, nsteps+1))

        return -1

sol = Solution()
result = sol.nearestExit([["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], [1,2])
print(result)