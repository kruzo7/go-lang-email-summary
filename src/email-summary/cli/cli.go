package cli

import (
	"flag"
	"os"

	"github.com/kruzo7/go-lang-email-summary/email-summary/reader"
)

func Run() {

	file := flag.String("file", "", "File to read from ex. -file=data.csv")
	output := flag.String("output", "console", "Generate output to 'console' (default) or to 'file' ex. -output=file")

	flag.Parse()

	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}

	reader.ReadFile(*file, *output)

}
