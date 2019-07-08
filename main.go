package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// App definition
	app := kingpin.New("gosra", "A command-line interface to get data from Sequence Read Archive")

	// Define help
	app.HelpFlag.Short('h')

	// Define get
	get(app)

	// Define search
	search(app)

	// parse and callback
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
