package main

type Accession struct {
	accid        string
	bioprojectid string
	biosampleid  []string
	experimentid []string
	runid        []string
}

func GetAccessions(id *string) (accid, expid, runid string) {
	return "DRA000001", "DRX000001", "DRR000001"
}
