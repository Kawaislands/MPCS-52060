package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Account represents a bank account with a balance
type Account struct {
	balance int64
}

// Transaction represents a bank transaction that moves
// a certain amount from one account to another account
type Transaction struct {
	From   *Account
	To     *Account
	amount int64
}

// SharedContext represents a wrapper object that holds
// all the shared resources of the goroutines between
// each other.
type SharedContext struct {
	transactions []*Transaction
	group        *sync.WaitGroup
}

// transfer executes a bank transaction
func transfer(transaction *Transaction) {

	from := transaction.From
	to := transaction.To
	amount := transaction.amount

	//Perform the transfer
	//Note: We are assuming the from.balance >= amount!
	atomic.AddInt64(&(from.balance), -amount)
	atomic.AddInt64(&(to.balance), amount)

}

// transferTask computes concurrent task where each goroutine
// will execute a certain number of transactions.
func transferTask(start, end int, context *SharedContext) {

	for i := start; i < end; i++ {
		//Retrieve transaction information
		transaction := context.transactions[i]
		transfer(transaction)
	}
	context.group.Done()

}

// forkJoin spawns off a certain number of goroutines to handle
// executing certain number of transactions.
func forkJoin(context *SharedContext) {

	/*
	  For this example of the fork-join pattern, we will hardcode in the threshold
	  and how many threads to correctly spawn based on the number of
	  transactions being less than or equal to 100. However, you will need to think
	  how you can do this dynamically.
	*/
	threshold := 100

	/* Sequential condition to perform the computation sequentially.
	 * Note: this is very specific trivial case here for only this specific example.
	 * In general you should basis this off some sequential threshold value.
	 */
	if len(context.transactions) <= threshold {
		for i := 0; i < threshold; i++ {
			transaction := context.transactions[i]
			from := transaction.From
			to := transaction.To
			amount := transaction.amount
			//Perform the transfer
			//Note: We are assuming the from.balance >= amount!
			from.balance -= amount
			to.balance += amount
		}
	} else {
		// Run the parallel code.
		numOfThreads := 5
		//This represents the granularity for each thread
		//Note: This example assumes each thread receives an equal number
		//of tasks. This will not always be the case.
		taskSize := len(context.transactions) / numOfThreads
		var start, end int
		//Spawn off a certain number of goroutines to execute
		//a specific number of transactions based on taskSize
		for thread := 0; thread < numOfThreads-1; thread++ {
			end = start + taskSize
			context.group.Add(1)
			go transferTask(start, end, context)
			start = end
		}
		context.group.Add(1)
		//Have the main goroutine handle the last portion.
		transferTask(start, end+taskSize, context)
		context.group.Wait()
	}
}
func main() {

	//Create accounts
	account1 := &Account{100}
	account2 := &Account{100}

	//Initialize all transactions
	transactions := make([]*Transaction, 0)

	for i := 0; i < 100; i++ {
		transactions = append(transactions, &Transaction{account1, account2, 1})
		transactions = append(transactions, &Transaction{account2, account1, 1})
	}

	/* Create a the shared resource value to hold information that will be shared
	 * between the goroutines
	 */
	group := &sync.WaitGroup{}
	context := &SharedContext{transactions, group}
	forkJoin(context)

	fmt.Println("Account1 Balance = $", account1.balance)
	fmt.Println("Account2 Balance = $", account2.balance)
}
