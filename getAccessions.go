package main

type Accession struct {
	accid        string
	bioprojectid string
	biosampleid  []string
	experimentid []string
	runid        []string
}

func GetAccessions(id *string) (expid, runid string) {
	return "DRX000001", "DRR000001"
}
