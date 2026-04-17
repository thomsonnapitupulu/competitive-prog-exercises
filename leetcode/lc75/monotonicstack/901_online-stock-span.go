package lc75

type StockSpanner struct {
	stack      []int
	nDaysStack []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

/*
*
Approach:
The StockSpanner structure utilizes two stacks: one for storing the prices (stack)
and another for storing the corresponding spans (ansStack). For each new price,
the method Next iteratively compares it with the top of the stack.
If the current price is greater, it pops from both stacks, summing up the spans.
This process repeats until we find a price on the top of the stack
that is greater than the current price, or the stack is empty.
The span for the current price is then calculated and stored, along with the price itself.
*
*/
func (this *StockSpanner) Next(price int) int {
	res := 1
	for len(this.stack) && price > this.stack[len(this.stack)-1] {
		this.stack = this.stack[:len(this.stack)-1] //pop
		res += this.nDaysStack[len(this.nDaysStack)-1]
		this.nDaysStack = this.nDaysStack[:len(this.nDaysStack)-1]
	}

	this.stack = append(this.stack, price)
	this.nDaysStack = append(this.nDaysStack, res)

	return res
}
