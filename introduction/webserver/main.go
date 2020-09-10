package main

import (
	"fmt"
	"sort"
)

func main() {
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello wolrd!"))
	// })

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	els := []int{9, 8, 7, 6, 5, 11}
	sort.Ints(els)
	fmt.Println(els)
}
