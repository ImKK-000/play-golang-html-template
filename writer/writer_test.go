package writer_test

import (
	. "play-golang/writer"
	"testing"
)

func Test_Writer_Should_Be_Data_Hello(t *testing.T) {
	expectedWriterData := "Hello"
	writerData := []byte("Hello")

	writer := Writer{}
	writer.Write(writerData)
	actualWriterData := string(writer.Data)

	if expectedWriterData != actualWriterData {
		t.Errorf("expect %s but it got %s", expectedWriterData, actualWriterData)
	}
}
