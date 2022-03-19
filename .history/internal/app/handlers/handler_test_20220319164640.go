package handlers

import (
	"net/http"
	"testing"
)

func TestModels_HandlerGet(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		m    Models
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.HandlerGet(tt.args.w, tt.args.r)
		})
	}
}
