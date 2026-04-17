package lc75

type MinHeap []int

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i,j int) bool {return h[i]<h[j]}
func (h MinHeap) Swap(i,j int) {h[i],h[j]=h[j],h[i]}
func (h *MinHeap) Push(x int){
    *h=append(*h,x)
    h.SiftUp()
}
func (h *MinHeap) Pop() int{
    lastElementIdx := h.Len()-1
    val := (*h)[0]
    h.Swap(0, lastElementIdx)
    *h = (*h)[:lastElementIdx]

    if h.Len() > 0 {
        h.SiftDown(0, h.Len()-1)
    }

    return val
}

func (h *MinHeap) SiftUp() {
    n:=h.Len()
    lastElementIdx:=n-1
    parentIdx:= (lastElementIdx - 1)/2

    for lastElementIdx > 0 && (*h)[lastElementIdx] < (*h)[parentIdx] {
        h.Swap(lastElementIdx, parentIdx)
        lastElementIdx = parentIdx
        parentIdx = (lastElementIdx - 1)/2
    }
}

func (h *MinHeap) SiftDown(startIdx, endIdx int) {
    leftChild:= (2*startIdx) + 1
    for leftChild <= endIdx {
        rightChild:= (2*startIdx) + 2

        if rightChild > endIdx {
            rightChild = -1
        }

        childIdx:=leftChild
        if rightChild != -1 && (*h)[leftChild] > (*h)[rightChild] {
            childIdx = rightChild
        }

        if (*h)[childIdx] < (*h)[startIdx] {
            h.Swap(startIdx, childIdx)
            startIdx = childIdx
            leftChild = (2*startIdx) + 1
        } else {
            return
        }
    }
}

type SmallestInfiniteSet struct {
    // this will contain the integers that were removed, and then added back
    // because we don't want to put all the integers (constraint say upto 1000)
    // this is more closer to the actual requirement that says this set will contain all positive integers
    nums MinHeap

    // this var will track the smallest number that can be popped.
    // if the nums have elements, they will be preferred, then we'll advance this counter 
    // when some one asks us to pop smallest
    current int

    // since we are not keeping / or in other words unable to put everything in the heap,
    // this optimization will help us not to look into the heap to find the number, because
    // the constraints says we want to have a set, if 2 same nums are added, only one gets added
    // so this map will keep track if the key is present
    numsTracker map[int]bool
}


func Constructor() SmallestInfiniteSet {
    return SmallestInfiniteSet {
        // we don't need to initialize nums
        current: 1,
        numsTracker: make(map[int]bool),
    }
}


func (this *SmallestInfiniteSet) PopSmallest() int {
   var poppedValue int
   if this.nums.Len() > 0 {
        // this means that the heap is not empty, the root will contain the minimum element
        poppedValue = this.nums.Pop()
        // also we need to mark this num in the set that now it is not available
        delete(this.numsTracker, poppedValue)
   } else {
        poppedValue = this.current
        this.current++
   }
   return poppedValue
}


func (this *SmallestInfiniteSet) AddBack(num int)  {
    // we already assume all numbers in front of current are already in the set
    if num >= this.current || this.numsTracker[num] {
        return
    }

    this.nums.Push(num)
    this.numsTracker[num]=true
}


/* === Ordered Set Implementation ===
*/

package main

type SmallestInfiniteSet struct {
	deleted map[int]struct{}
	current int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		deleted: make(map[int]struct{}),
		current: 1,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	c := this.current
	for _, ok := this.deleted[c]; ok; _, ok = this.deleted[c] {
		c++
	}
	this.deleted[c] = struct{}{}
	this.current = c + 1
	return c
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	delete(this.deleted, num)
	if num < this.current {
		this.current = num
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

