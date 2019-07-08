package main

import "testing"

func TestGetAccessions(t *testing.T) {
	tests := []struct {
		name              string
		id, acc, exp, run string
	}{
		{"with RunID", "DRR000001", "DRA000001", "DRX000001", "DRR000001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out1, out2, out3 := GetAccessions(&tt.id)
			if out1 != tt.acc {
				t.Errorf("failed GetAccessions()")
			}
			if out2 != tt.exp {
				t.Errorf("failed GetAccessions()")
			}
			if out3 != tt.run {
				t.Errorf("failed GetAccessions()")
			}

		})
	}

}
