package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
)

func getRating(wg *sync.WaitGroup, totalRating *int){
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	sleepRand := rand.Intn(10)
	time.Sleep(time.Duration(sleepRand)*time.Millisecond)
	mutex.Lock()
	score := rand.Intn(10)
	*totalRating+=score
	mutex.Unlock()
}


func main(){
	totalRating := 0
	var wg sync.WaitGroup
	for i:=0;i<200;i++{
		wg.Add(1)
		go getRating(&wg,&totalRating)
	}
	wg.Wait()
	ans := float64(totalRating)/200
	fmt.Println(ans)
}
