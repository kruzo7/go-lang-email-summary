package reader

import (
	"os"

	"github.com/kruzo7/go-lang-email-summary/email-summary/internal"
	"github.com/kruzo7/go-lang-email-summary/email-summary/parser"
)

func ReadFile(file string, output string) {

	ifFileExists(file)

	parser.Parse(file, output)
}

func ifFileExists(file string) {

	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		internal.Trace(err, true)
	}
}
