package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Result struct {
	UtcTime time.Time
	Time    time.Time
	TZ      string
	Offset  int
}

func handler(w http.ResponseWriter, r *http.Request) {
	// check for request Header and return it
	val, ok := r.Header["x-request-id"]
	if ok {
		w.Header().Add("x-request-id", val[0])
	}

	dt := time.Now()

	zone_name, offset := dt.Zone()

	res := Result{
		Time:    dt,
		UtcTime: dt.UTC(),
		TZ:      zone_name,
		Offset:  offset,
	}
	str, _ := json.Marshal(res)
	fmt.Fprintf(w, string(str))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
