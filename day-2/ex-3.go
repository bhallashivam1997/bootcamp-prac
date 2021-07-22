package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
)

func deposit(balance *float64, amount float64){
	mutex.Lock()
	*balance+=amount
	fmt.Println("Deposit Complete")
	mutex.Unlock()
}

func withdraw(balance *float64, amount float64){
	mutex.Lock()
	if *balance<amount{
		fmt.Println("Insufficient Balance !!")
	}else{
		*balance-=amount
		fmt.Println("Withdrawal Complete")
	}
	mutex.Unlock()
}

func main(){
	balance := 500.0
	go withdraw(&balance,1000)
	time.Sleep(5*time.Microsecond)
	go deposit(&balance,500)
	time.Sleep(5*time.Microsecond)
	go withdraw(&balance,1000)
}
