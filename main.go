package main

import (
	// "encoding/json"
	"fmt"
	"github.com/couchbaselabs/go-couchbase"
	"ilpaijin"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	event := getEvent(dbConnect())
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<code>%v</code>", event)
}

func dbConnect() *couchbase.Bucket {
	c, err := couchbase.Connect("http://localhost:8091")

	if err != nil {
		fmt.Printf("Errore connessione couchbase, %v", err)
	}

	pool, err := c.GetPool("default")

	bt, err := pool.GetBucket("livebetting")

	if err != nil {
		fmt.Printf("Errore bucket, %v", err)
	}

	return bt
}

func getEvent(b *couchbase.Bucket) map[string]interface{} {
	var e map[string]interface{}

	b.Get("Event::111220", &e)

	return e
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9000", nil)

	if x := ilpaijin.Myname(); x == "Paolo" {
		fmt.Println("\nciao", x)
	} else {
		fmt.Println("\nciao sconosciuto")
	}

	// jsonObj := []byte(`{"0": "a", "1": "b", "2": {"0": "c"}}`)
	// var response_hash map[string]interface{}
	// _ = json.Unmarshal(jsonObj, &response_hash)

	// fmt.Println(response_hash)

	// for k, v := range response_hash {
	// 	fmt.Println("k", k, "vale", v)
	// }
}
