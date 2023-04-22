package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		msg := fmt.Sprintf("type is not http.Handler, but is %T", v)
		t.Error(msg)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		msg := fmt.Sprintf("type is not http.Handler, but is %T", v)
		t.Error(msg)
	}
}
