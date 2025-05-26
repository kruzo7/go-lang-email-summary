package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/kruzo7/go-lang-email-summary/email-summary/formatter"
	"github.com/kruzo7/go-lang-email-summary/email-summary/internal"
)

const comma = ','
const invalidkey = "invalid"

func Parse(filename string, output string) {

	file, err := os.Open(filename)
	internal.Trace(err, true)
	defer file.Close()

	csvreader := csv.NewReader(file)
	csvreader.Comma = comma

	readheader(csvreader)

	f := formatter.New()

	for {

		record, err := csvreader.Read()
		internal.Trace(err, false)
		if checkeof(err) {
			break
		}

		domain := parsedomain(record)
		f.Add(domain)
	}

	f.Print(output)
}

func readheader(reader *csv.Reader) {
	if _, err := reader.Read(); err != nil {
		internal.Trace(err, true)
	}
}

func parsedomain(record []string) string {
	email := record[2]

	emailparts := strings.Split(email, "@")

	if len(emailparts) != 2 {
		internal.Trace(fmt.Errorf("invalid email address: %s", email), false)
		return invalidkey
	}

	return emailparts[1]
}

func checkeof(err error) bool {
	return err != nil && err.Error() == "EOF"
}
