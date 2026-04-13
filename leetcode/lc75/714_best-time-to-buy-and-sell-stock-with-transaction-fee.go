package lc75

func maxProfit(prices []int, fee int) int {
	holdCash := 0           //stands for hold cash on day 0
	holdStock := -prices[0] //stands for hold stock on day 0

	sellStock, buyStock := 0, 0
	for i := 1; i < len(prices); i++ {
		//if the decision is to hold cash on day i, there are 2 options
		//option 1: sell the stock held from previous day
		//option 2: do nothing, keep on holding cash like previous day
		//max profit of day i with cash is max(option1, option2)
		sellStock = prices[i] + holdStock - fee
		holdCash = max(sellStock, holdCash)

		//if the decision is to hold stock on day i, there are also 2 options
		//option 1: buy stock, based on previous day's cash profit
		//option 2: do nothing, keep on holding stock like previous day
		//max profit of day i with stock is max(option1, option2)
		buyStock = holdCash - prices[i]
		holdStock = max(buyStock, holdStock)
	}

	return holdCash
}
