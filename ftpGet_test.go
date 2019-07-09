package main

import (
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func TestFtpGet(t *testing.T) {
	tests := []struct {
		name, url string
	}{
		{"DRR007230 from EBI", "ftp://ftp.sra.ebi.ac.uk/vol1/drr/DRR007/DRR007230"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.url)
			FtpGet(u)
			fileName := filepath.Base(tt.url) + ".sra"
			if exists(fileName) != true {
				t.Errorf("Failed FtpGet()")
			}
		})
	}

}
