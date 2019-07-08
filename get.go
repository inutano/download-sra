package main

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

type Accession struct {
	accid        string
	bioprojectid string
	biosampleid  []string
	experimentid []string
	runid        []string
}

func getAccessions(id *string) (accid, expid, runid string) {
	return "DRA000001", "DRX000001", "DRR000001"
}

func getEBIURL(runid string) (sraUrl *url.URL) {
	baseUrl := "ftp://ftp.sra.ebi.ac.uk/vol1/"
	sraUrl, _ = url.Parse(baseUrl)
	sraUrl.Path = path.Join(sraUrl.Path, strings.ToLower(runid[0:3]), runid[0:6])

	switch len(runid) {
	case 9:
		sraUrl.Path = path.Join(sraUrl.Path, runid)
	case 10:
		sraUrl.Path = path.Join(sraUrl.Path, "00", runid[len(runid)-1:], runid)
	default:
	}
	return
}

func getNCBIURL(runid string) (sraUrl *url.URL) {
	baseUrl := "ftp://ftp.ncbi.nlm.nih.gov/sra/sra-instant/reads/ByRun/sra"
	sraUrl, _ = url.Parse(baseUrl)
	sraUrl.Path = path.Join(sraUrl.Path, runid[0:3], runid[0:6], runid, runid+".sra")
	return
}

func getDDBJURL(runid, expid string) (sraUrl *url.URL) {
	baseUrl := "ftp://ftp.ddbj.nig.ac.jp/ddbj_database/dra/sralite/ByExp/litesra"
	sraUrl, _ = url.Parse(baseUrl)
	sraUrl.Path = path.Join(sraUrl.Path, expid[0:3], expid[0:6], expid, runid, runid+".sra")
	return
}

func getUrl(repo *string, accid, expid, runid string) (sraUrl *url.URL) {
	switch *repo {
	case "ebi":
		sraUrl = getEBIURL(runid)
	case "ncbi":
		sraUrl = getNCBIURL(runid)
	case "ddbj":
		sraUrl = getDDBJURL(runid, expid)
	}
	return
}

func get(app *kingpin.Application) {
	cmd := app.Command("get", "get sequence data")
	id := cmd.Arg("id", "ID to get data").Required().String()
	repo := cmd.Flag("repo", "Repository to use").Default("ebi").String()

	cmd.Action(func(c *kingpin.ParseContext) error {
		accid, expid, runid := getAccessions(id)
		url := getUrl(repo, accid, expid, runid)

		switch *repo {
		case "ebi":
			fmt.Printf("get %s from EBI, UK\n", url)
		case "ncbi":
			fmt.Printf("get %s from NCBI, US\n", url)
		case "ddbj":
			fmt.Printf("get %s from DDBJ, Japan\n", url)
		}
		return nil
	})
}
