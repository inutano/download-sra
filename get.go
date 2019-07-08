package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

func get(app *kingpin.Application) {
	cmd := app.Command("get", "get sequence data")
	id := cmd.Arg("id", "ID to get data").Required().String()
	repo := cmd.Flag("repo", "Repository to use").Default("ebi").String()

	cmd.Action(func(c *kingpin.ParseContext) error {
		accid, expid, runid := GetAccessions(id)
		url := GetUrl(repo, accid, expid, runid)

		fmt.Printf("Downloading data from %s\n", url)
		FtpGet(url)
		fmt.Printf("Download finished.\n")

		return nil
	})
}
