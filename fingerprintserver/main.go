package main

import (
	"encoding/json"
	"fmt"
	fingerprints "github.com/SpectralHiss/code-test/fingerprintserver/eventreport"
	"log"
	"net/http"
	//"io/ioutil"
)

type Reports map[string]fingerprints.Data

func main() {
	reports := Reports{}

	http.HandleFunc("/result/resize", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("access-control-allow-headers", "content-type, accept")

		resize := fingerprints.ResizeEventReport{}

		err := json.NewDecoder(req.Body).Decode(&resize)
		if err != nil {
			return
		}

		report, ok := reports[resize.SessionID]

		if !ok {
			report = initialiseReport(reports, resize.SessionID, resize.WebsiteURL)
		}

		report.ResizeFrom = resize.ResizeFrom
		report.ResizeTo = resize.ResizeTo

		reports[resize.SessionID] = report

		fmt.Printf("partial report for user session %s:\n", resize.SessionID)
		printReport(reports[resize.SessionID])

	})

	http.HandleFunc("/result/copypaste", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("access-control-allow-headers", "content-type, accept")

		copypaste := fingerprints.CopyEventReport{}

		err := json.NewDecoder(req.Body).Decode(&copypaste)
		if err != nil {
			return
		}

		report, ok := reports[copypaste.SessionID]

		if !ok {
			report = initialiseReport(reports, copypaste.SessionID, copypaste.WebsiteURL)
		}

		if report.CopyAndPaste == nil {
			report.CopyAndPaste = map[string]bool{
				copypaste.FormID: copypaste.Pasted,
			}
		} else {
			report.CopyAndPaste[copypaste.FormID] = copypaste.Pasted
		}

		reports[copypaste.SessionID] = report

		fmt.Printf("partial report for user session %s:\n", copypaste.SessionID)

		printReport(reports[copypaste.SessionID])

	})

	http.HandleFunc("/result/delay", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("access-control-allow-headers", "content-type, accept")

		delay := fingerprints.DelayEventReport{}

		err := json.NewDecoder(req.Body).Decode(&delay)
		if err != nil {
			return
		}

		report, ok := reports[delay.SessionID]

		if !ok {
			report = initialiseReport(reports, delay.SessionID, delay.WebsiteURL)
		}

		report.FormCompletionTime = delay.Time

		reports[delay.SessionID] = report

		fmt.Printf("\n\n Final report for user session %s \n\n", delay.SessionID)
		printReport(reports[delay.SessionID])

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initialiseReport(reports map[string]fingerprints.Data, sessionID, websiteURL string) fingerprints.Data {
	fmt.Printf("\n a new user session has been started %s:\n", sessionID)
	reports[sessionID] = fingerprints.Data{
		WebsiteUrl: websiteURL,
		SessionId:  sessionID,
	}
	return reports[sessionID]
}

func printReport(report fingerprints.Data) {
	bytes, err := json.Marshal(report)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", string(bytes))
}
