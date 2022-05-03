package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	wg := &sync.WaitGroup{} // WaitGroupの値を作る

	for i := 0; i < 10; i++ {
		log.Println("i:", i)
		wg.Add(1)
		go func() {
			log.Println("i:", i)
			log.Println("path:", r.URL.Path)
			wg.Done()
		}()
	}
	log.Println("finish before wait")
	fmt.Fprintf(w, "Hello astaxie!")
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