package main

import (
	"net/url"
	"path"
	"strings"
)

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

func GetUrl(repo *string, expid, runid string) (sraUrl *url.URL) {
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
