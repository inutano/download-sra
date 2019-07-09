package main

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"
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
		{"DRR007230 from DDBJ", "ftp://ftp.ddbj.nig.ac.jp/ddbj_database/dra/sralite/ByExp/litesra/DRX/DRX006/DRX006419/DRR007230/DRR007230.sra"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, _ := url.Parse(tt.url)
			FtpGet(u)
			fileName := strings.TrimSuffix(filepath.Base(tt.url), filepath.Ext(tt.url)) + ".sra"
			if exists(fileName) != true {
				t.Errorf("Failed FtpGet()")
			}
			if err := os.Remove(fileName); err != nil {
				t.Errorf("Failed FtpGet()")
			}
		})
	}

}
