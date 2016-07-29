package main

import (
	"encoding/json"
	"fmt"
	event "github.com/SpectralHiss/code-test/fingerprintserver/eventreport"
	"net"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestStartedOK(t *testing.T) {

	cmd := exec.Command("fingerprintserver")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		t.Error("server failed to start", err)
	}

	time.Sleep(2 * time.Second)

	if isOpen(8080) {
		t.Error("server is not listening")
	}

	client := &http.Client{}

	// send resize
	f, err := os.Open("./fixtures/resize.json")
	if err != nil {
		panic(err)
	}

	output := make(chan event.ResizeEventReport)

	go func(c chan event.ResizeEventReport) {
		var result = event.ResizeEventReport{}
		if err := json.NewDecoder(os.Stdin).Decode(&result); err != nil {
			t.Error("we do not see the object being created in terminal")
		}
		c <- result
	}(output)

	// just checking for resize the rest can be manually tested through
	// the client web app

	_, err = client.Post("http://localhost:8080/resize", "application/json", f)
	if err != nil {
		t.Error("request to server failed", err)
	}

	testResult := <-output

	expected := event.Data{
	WebsiteUrl         "something"
	SessionId          "something"
	ResizeFrom         event.Dimension{"height":"something", "width":"something"}
	ResizeTo event.Dimension{"height":"something", "width":"something"}
	}

	if result != expected {
		t.Error(result ," does not match ", expected);
	} 

	if err := cmd.Wait(); err != nil {
		t.Fatal(err)
	}

}

// helpers:
func isOpen(port int) bool {

	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "localhost", port))
	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}
