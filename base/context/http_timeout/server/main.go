package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	intn := rand.Intn(2)
	if intn == 0 {
		time.Sleep(10 * time.Second)
		_, _ = fmt.Fprintf(w, "slow")
		return
	}
	_, _ = fmt.Fprintf(w, "quick")
}
func main() {
	http.HandleFunc("/", indexHandle)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
