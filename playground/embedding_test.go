package main

import (
	"bytes"
	"log"
	"testing"
)

func TestReadWriterStructImplementsReadWriter(t *testing.T) {
	var buf bytes.Buffer
	var rw ReadWriter = NewReadWriterStruct(&buf, &buf) // ReadWriter 인터페이스 타입으로 선언

	inputText := "test input"
	if _, err := rw.Write([]byte(inputText)); err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// bufio.Writer의 Flush를 호출해야 실제로 버퍼에 쓰기 작업이 완료됨
	if w, ok := rw.(*ReadWriterStruct); ok {
		w.Writer.Flush()
	} else {
		t.Fatalf("rw is not of type *ReadWriterStruct")
	}

	output := make([]byte, len(inputText))
	if _, err := rw.Read(output); err != nil {
		t.Fatalf("Read failed: %v", err)
	}

	if string(output) != inputText {
		t.Errorf("Expected '%s', got '%s'", inputText, string(output))
	}
}

func TestReadWriterStruct(t *testing.T) {
	var buf bytes.Buffer
	rw := NewReadWriterStruct(&buf, &buf)

	input := "hello, world"
	rw.Writer.Write([]byte(input))
	rw.Writer.Flush()

	output := make([]byte, len(input))
	rw.Reader.Read(output)

	if string(output) != input {
		t.Errorf("expected %q, got %q", input, string(output))
	}
}

func TestJobLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "Job: ", log.Ldate)
	job := NewJob("Test Command", logger)

	job.Println("starting now...")

	if !bytes.Contains(buf.Bytes(), []byte("starting now...")) {
		t.Error("log does not contain the expected message")
	}
}
