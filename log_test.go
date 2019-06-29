package loggy

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
	"testing"
)

func TestNotice(t *testing.T) {
	sut := NewLoggy()
	writer := bytes.NewBuffer(make([]byte, 0))
	sut.SetWriter(writer)

	expected := "hello world"
	sut.Notice(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestWarn(t *testing.T) {
	sut := NewLoggy()
	writer := bytes.NewBuffer(make([]byte, 0))
	sut.SetWriter(writer)
	expected := "hello world"
	sut.Warn(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestError(t *testing.T) {
	sut := NewLoggy()
	writer := bytes.NewBuffer(make([]byte, 0))
	sut.SetWriter(writer)

	expected := "hello world"
	sut.Error(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log error doesn't match expected")
	}

	writer = bytes.NewBuffer(make([]byte, 0))
	sut.SetWriter(writer)
	err := errors.New(expected)
	sut.Error(err)
	actual = writer.String()
	if expected != actual {
		log.Println(actual)
		t.Error("log error doesn't match expected")
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if e := recover(); e == nil {
			t.Error("expected a panic, that was not received")
		}
	}()
	sut := NewLoggy()
	writer := bytes.NewBuffer(make([]byte, 0))
	sut.SetWriter(writer)

	expected := "hello world"
	sut.Panic(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestEnablePretty(t *testing.T) {
	sut := NewLoggy()
	sut.EnablePretty()
	writer := bytes.NewBuffer(make([]byte, 0))
	sut.SetWriter(writer)

	colour := aurora.NewAurora(true)
	expected := "hello world"
	sut.Error(expected)
	actual := writer.String()
	expected = fmt.Sprint(colour.Red(expected))
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestDisablePretty(t *testing.T) {
	sut := NewLoggy()
	sut.EnablePretty()
	sut.DisablePretty()

	writer := bytes.NewBuffer(make([]byte, 0))
	sut.SetWriter(writer)

	expected := "hello world"
	sut.Error(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}

func TestSetDestination(t *testing.T) {
	sut := NewLoggy()
	writer := bytes.NewBuffer(make([]byte, 0))
	sut.SetWriter(writer)

	expected := "hello world"
	sut.Notice(expected)
	actual := writer.String()
	if expected != actual {
		t.Error("log info doesn't match expected")
	}
}
