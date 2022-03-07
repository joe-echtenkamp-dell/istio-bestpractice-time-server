package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Result struct {
	UtcTime time.Time
	Time    time.Time
	TZ      string
	Offset  int
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("time server recieved request")
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
	fmt.Fprint(w, string(str))
}

func main() {
	log.SetOutput(os.Stdout)
	log.Print("Starting time server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
