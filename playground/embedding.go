package main

import (
	"bufio"
	"io"
	"log"
)

type ReadWriter interface {
	io.Reader
	io.Writer
}

type ReadWriterStruct struct {
	*bufio.Reader
	*bufio.Writer
}

func NewReadWriterStruct(reader io.Reader, writer io.Writer) *ReadWriterStruct {
	return &ReadWriterStruct{
		Reader: bufio.NewReader(reader),
		Writer: bufio.NewWriter(writer),
	}
}

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{Command: command, Logger: logger}
}
