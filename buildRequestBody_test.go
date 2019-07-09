package main

import (
	"testing"
)

func TestBuildRequestBody(t *testing.T) {
	tests := []struct {
		name, keyword, date, reads, bases string
	}{
		{"Simple", "Ohanami", "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := BuildRequestBody(tt.keyword, tt.date, tt.reads, tt.bases)

			if len(buf.String()) < 1 {
				t.Errorf("Failed BuildRequestBody()")
			}
		})
	}
}
