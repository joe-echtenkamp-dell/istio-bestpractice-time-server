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
	log.Println("time server recieved request")
	// check for request Header and return it
	header := http.CanonicalHeaderKey("x-request-id")
	val, ok := r.Header[header]
	if ok {
		w.Header().Add(header, val[0])
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
	// log.SetOutput(os.Stdout)
	log.Println("Starting time server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
