package main

import "testing"

func TestGetUrl(t *testing.T) {
	tests := []struct {
		repo, exp, run, want string
	}{
		{"ddbj", "DRX000001", "DRR000001", "ftp://ftp.ddbj.nig.ac.jp/ddbj_database/dra/sralite/ByExp/litesra/DRX/DRX000/DRX000001/DRR000001/DRR000001.sra"},
		{"ebi", "DRX000001", "DRR000001", "ftp://ftp.sra.ebi.ac.uk/vol1/drr/DRR000/DRR000001"},
		{"ebi", "DRX0000001", "DRR0000001", "ftp://ftp.sra.ebi.ac.uk/vol1/drr/DRR000/001/DRR0000001"},
		{"ncbi", "DRX000001", "DRR000001", "ftp://ftp.ncbi.nlm.nih.gov/sra/sra-instant/reads/ByRun/sra/DRR/DRR000/DRR000001/DRR000001.sra"},
	}
	for _, tt := range tests {
		t.Run(tt.repo, func(t *testing.T) {
			if got := GetUrl(&tt.repo, tt.exp, tt.run); got.String() != tt.want {
				t.Fatalf("want = %s, got = %s", tt.want, got)
			}
		})
	}
}
