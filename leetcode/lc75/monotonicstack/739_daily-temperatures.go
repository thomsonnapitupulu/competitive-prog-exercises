package main

func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[index] = i - index
		}

		stack = append(stack, i)
	}

	return results
}

// i=0
//     stack=[0]
// i=1
//     index=0
//     stack=[]
//     answer=[1,0,0,0,0,0,0]
//     stack=[1]
// i=2
//     index=1
//     stack=[]
//     answer=[1,1,0,0,0,0,0]
//     stack=[2]
// i=3
//     stack=[2,3]
// i=4
//     stack=[2,3,4]
// i=5
//     stack=[2,3,4]
//     index=4
//     stack=[2,3]
//     answer=[1,1,0,0,1,0,0]
//     index=3
//     stack=[2]
//     answer=[1,1,2,0,1,0,0]
//     stack=[2,5]
// i=6
//     index=5
//     stack=[2]
//     answer=[1,1,2,0,1,1,0]
//     index=2
//     stack=[]
//     answer=[1,1,2,4,1,1,0]
