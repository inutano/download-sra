package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

func get(app *kingpin.Application) {
	// Define get option
	cmd := app.Command("get", "get sequence data")

	// Parse arguments
	repo := cmd.Flag("repo", "Repository to use").Default("ebi").String()
	id := cmd.Arg("id", "ID to get data").Required().String()

	cmd.Action(func(c *kingpin.ParseContext) error {
		switch *repo {
		case "ebi":
			fmt.Printf("get %s from EBI, UK", *id)
		case "ncbi":
			fmt.Printf("get %s from NCBI, US", *id)
		case "ddbj":
			fmt.Printf("get %s from DDBJ, Japan", *id)
		}
		return nil
	})
}
