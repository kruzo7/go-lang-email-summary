package formatter

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kruzo7/go-lang-email-summary/email-summary/internal"
)

type Formatter struct {
	emailsnumber map[string]int
}

func New() Formatter {
	f := Formatter{
		emailsnumber: make(map[string]int),
	}

	return f
}

func (f *Formatter) Add(key string) {
	if f.searchemail(key) {
		f.increment(key)
	} else {
		f.emailsnumber[key] = 1
	}
}

func (f *Formatter) Print(output string) string {

	sorted := f.SortEmailsByValueDesc()
	filename := ""

	switch output {
	case "console":
		f.outputconsole(sorted)
	case "file":
		filename = f.outputfile(sorted)
	default:
		f.outputconsole(sorted)
	}

	return filename
}

func (f *Formatter) searchemail(key string) bool {
	_, exists := f.emailsnumber[key]
	return exists
}

func (f *Formatter) increment(key string) {
	f.emailsnumber[key]++

}

func (f *Formatter) outputconsole(sorted []Pair) {

	for _, pair := range sorted {
		fmt.Println(pair.Key, pair.Value)
	}
}

func (f *Formatter) outputfile(sorted []Pair) string {

	filename := strconv.FormatInt(time.Now().Unix(), 10) + ".txt"

	ofile, err := os.Create(filename)
	internal.Trace(err, true)
	defer ofile.Close()

	for _, pair := range sorted {
		_, err := ofile.WriteString(fmt.Sprintf("%s %d\n", pair.Key, pair.Value))
		internal.Trace(err, true)
	}

	ofile.Sync()

	return filename
}
