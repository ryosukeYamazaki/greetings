package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{} // WaitGroupの値を作る
	log.Println("path:", r.URL.Path)
	concurrentNum := 10
	tmpNum := make(chan int)

	for i := 0; i < concurrentNum; i++ {
		log.Println("i (sync):", i)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println("i (async):", i)
			tmpNum <- i
			// 結果をこの辺で受け取る
		}(i)
	}
	log.Println("finish before wait")
	fmt.Fprintf(w, "Hello astaxie!")

	sumNum := 0
	for i := 0; i < concurrentNum; i++ {
		sumNum += <-tmpNum
		log.Println("sumNum: ", sumNum)
	}
	wg.Wait()
	log.Println("finish after wait")
}

func main() {
	http.HandleFunc("/hogehoge", sayhelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
