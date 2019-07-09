package main

import "gopkg.in/alecthomas/kingpin.v2"

func search(app *kingpin.Application) {
	cmd := app.Command("search", "search sequence data by keyword")

	// entry : select ES index to search
	entry := cmd.Arg("entry", "Search entry, biosample or experiment").String()

	// keyword, date, reads, bases : search conditions
	keyword := cmd.Arg("keyword", "Search query keyword").Required().String()
	date := cmd.Arg("date", "Filtering by date").String()
	reads := cmd.Arg("reads", "Filtering by reads").String()
	bases := cmd.Arg("bases", "Filtering by bases").String()

	// format, download : search options
	format := cmd.Arg("format", "Output result format (txt, tsv, json)").String()
	download := cmd.Flag("download", "Download data found by query").Bool()

	cmd.Action(func(c *kingpin.ParseContext) error {
		request := BuildRequestBody(*keyword, *date, *reads, *bases)
		result := SearchRequest(*entry, request)
		SearchOutput(result, *format)
		if *download == true {
			// download data using result
		}
		return nil
	})
}
