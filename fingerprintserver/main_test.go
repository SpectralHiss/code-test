package main

import (
	"encoding/json"
	"fmt"
	event "github.com/SpectralHiss/code-test/fingerprintserver/eventreport"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"testing"
	"time"
)

func TestStartedAndResizeOK(t *testing.T) {

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
	reqJSON, err := os.Open("./fixtures/resize_request.json")
	if err != nil {
		panic(err)
	}

	output := make(chan event.Data, 1)

	_, err = client.Post("http://localhost:8080/resize", "application/json", reqJSON)

	go func() {
		var result = event.Data{}
		if err := json.NewDecoder(os.Stdin).Decode(&result); err != nil {
			t.Error("we do not see the object being created in terminal")
		}

		output <- result
	}()

	// just checking for resize the rest can be manually tested through
	// the client web app

	if err != nil {

		t.Error("request to server failed", err)
	}

	select {
	case testResult := <-output:
		expected := event.Data{
			WebsiteUrl: "http://localhost:8080",
			SessionId:  "123123-123123-123123123",
			ResizeFrom: event.Dimension{Height: "391", Width: "400"},
			ResizeTo:   event.Dimension{Height: "640", Width: "1035"},
		}

		if !reflect.DeepEqual(testResult, expected) {
			t.Error(testResult, " does not match ", expected)
		}

	case <-time.After(time.Second * 10):
		t.Error("test timed out")
	}

	//cmd.Process.Kill()

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
