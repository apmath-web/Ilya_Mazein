package main

import (
	"net/http"
	"fmt"
	"time"
)

func datime(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, time.Now())
}

func waiter(w http.ResponseWriter, r *http.Request) {
	amount := r.URL.Query().Get("delay")+"ms"
	duration, _ := time.ParseDuration(amount)
	time.Sleep(duration)
	fmt.Fprintf(w, "Sleep for "+amount+" completed")
}

func main() {
	http.HandleFunc("/date", datime)
	http.HandleFunc("/wait", waiter)
	http.ListenAndServe(":8080", nil)
}
