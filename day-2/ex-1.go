package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
)


func countFreq(s string, wg *sync.WaitGroup, freq map[byte]int){
	mutex.Lock()
	for i:=0;i<len(s);i++{
		freq[s[i]]++
	}
	mutex.Unlock()
	wg.Done()
}

func main() {
	words := []string{"quick","brown","fox","lazy","dog"}
	freq := make(map[byte]int)
	var wg sync.WaitGroup

	for _,word :=range words{
		wg.Add(1)
		go countFreq(word,&wg,freq)
	}

	wg.Wait()
	ans :=make(map[string]int)
	for i:=0;i<26;i++{
		if val,ok :=freq[byte(i+'a')]; ok{
			ans[string(byte(i+'a'))] = val
		}else{
			ans[string(byte(i+'a'))] = 0
		}
	}
	stringJson,err := json.Marshal(ans)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(stringJson))
}
