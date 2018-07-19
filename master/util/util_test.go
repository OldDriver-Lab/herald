package util

import (
	"net/http"
	"testing"
)

func TestResponseOK(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		body interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResponseOK(tt.args.w, tt.args.body)
		})
	}
}

func TestResponseError(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResponseError(tt.args.w, tt.args.code, tt.args.message)
		})
	}
}
