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
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		t.Error("server failed to start", err)
	}

	time.Sleep(1 * time.Second)

	if isOpen(8080) {
		t.Error("server is not listening")
	}

	client := &http.Client{}

	// send resize
	f, err := os.Open("./fixtures/resize.json")
	if err != nil {
		panic(err)
	}

	_, err = client.Post("resize", "application/json", f)
	if err != nil {
		t.Error("request to server failed", err)
	}

	if err := json.NewDecoder(stdout).Decode(&event.ResizeEventReport{}); err != nil {
		t.Error("we do not see the object being created in terminal")
	}

	//fmt.Printf(stdout.)
	//cmd.Process.Kill()
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
