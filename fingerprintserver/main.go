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

	http.HandleFunc("/resize", func(resp http.ResponseWriter, req *http.Request) {
		resize := fingerprints.ResizeEventReport{}

		json.NewDecoder(req.Body).Decode(&resize)

		report, ok := reports[resize.SessionID]

		if !ok {
			report = initialiseReport(reports, resize.SessionID, resize.WebsiteURL)
		}

		report.ResizeFrom = resize.ResizeFrom
		report.ResizeTo = resize.ResizeTo

		reports[resize.SessionID] = report

		fmt.Printf("partial report for user session %s:", resize.SessionID)
		printReport(reports[resize.SessionID])

	})

	http.HandleFunc("/copypaste", func(resp http.ResponseWriter, req *http.Request) {
		copypaste := fingerprints.CopyEventReport{}

		json.NewDecoder(req.Body).Decode(&copypaste)

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

		fmt.Printf("partial report for user session %s:", copypaste.SessionID)

		printReport(reports[copypaste.SessionID])

	})

	http.HandleFunc("/delay", func(resp http.ResponseWriter, req *http.Request) {
		delay := fingerprints.DelayEventReport{}

		json.NewDecoder(req.Body).Decode(&delay)

		report, ok := reports[delay.SessionID]

		if !ok {
			report = initialiseReport(reports, delay.SessionID, delay.WebsiteURL)
		}

		report.FormCompletionTime = delay.Time

		reports[delay.SessionID] = report

		fmt.Printf("final report for user session %s:", delay.SessionID)
		printReport(reports[delay.SessionID])

	})

	//	http.HandleFunc("/copypaste", func(resp http.ResponseWriter, req *http.Request) {}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initialiseReport(reports map[string]fingerprints.Data, sessionID, websiteURL string) fingerprints.Data {
	fmt.Printf("a new user session has been started %s", sessionID)
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
