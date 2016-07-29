package main

import (
	"encoding/json"
	"fmt"
	fingerprints "github.com/SpectralHiss/code-test/fingerprintserver/eventreport"
	"log"
	"net/http"
)

func main() {
	//report := fingerprints.Data{}

	http.HandleFunc("/resize", func(resp http.ResponseWriter, req *http.Request) {
		resize := fingerprints.ResizeEventReport{}
		json.NewDecoder(req.Body).Decode(&resize)
		fmt.Printf("%#v", resize)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
