package command

import (
	"testing"
)

func TestPingMustWork(t *testing.T) {
	cmd := New(".", "ping 8.8.8.8 -c 1")

	_, err := cmd.Run()

	if err != nil {
		t.Errorf("ping wont have to fail")
	}
}

func TestPingMustFail(t *testing.T) {
	cmd := New(".", "ping 8.8.8.8.8 -c 1")

	_, err := cmd.Run()

	if err == nil {
		t.Errorf("ping have to fail")
	}
}

func TestShouldCaptureStdout(t *testing.T) {
	textToPrint := "gohotreloader"
	expectedOutput := textToPrint + "\n"
	cmd := New(".", "echo "+textToPrint)

	output, err := cmd.Run()

	if err != nil {
		t.Errorf("ping wont have to fail")
	}
	if output != expectedOutput {
		t.Errorf("output is not captured correctly\nexpected: %s, but %s was found", expectedOutput, output)
	}
}

func TestShouldCaptureStderr(t *testing.T) {
	expectedOutput := "ping: 8.8.8.8.8: Name or service not known\n"
	cmd := New(".", "ping 8.8.8.8.8 -c 1")

	output, err := cmd.Run()

	if err == nil {
		t.Errorf("ping have to fail")
	}
	if output != expectedOutput {
		t.Errorf("output is not captured correctly\nexpected: %s, but %s was found", expectedOutput, output)
	}
}
