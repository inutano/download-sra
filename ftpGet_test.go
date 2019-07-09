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
		{"SRR1215802 from DDBJ", "ftp://ftp.ddbj.nig.ac.jp/ddbj_database/dra/sralite/ByExp/litesra/SRX/SRX509/SRX509580/SRR1215802/SRR1215802.sra"},
		{"SRR1215802 from NCBI", "ftp://ftp.ncbi.nlm.nih.gov/sra/sra-instant/reads/ByRun/sra/SRR/SRR121/SRR1215802/SRR1215802.sra"},
		{"SRR1215802 from EBI", "ftp://ftp.sra.ebi.ac.uk/vol1/srr/SRR121/002/SRR1215802"},
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
