package main

import (
	"os"
	"text/template"
	"fmt"
	"io/ioutil"
	"bytes"
)

const (
	journalFile = "/tmp/.touchfile"
)

var (
	journalMessages = make(map[string]bool)
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createFile(filename string, newData []byte, perm os.FileMode) (err error) {
	// this is just a vanity wrapper
	err = nil

	err = ioutil.WriteFile(filename, newData, perm)

	return

}

func updateFile(filename string, newData []byte, perm os.FileMode) (changed bool, err error) {
	changed = false
	err = nil
	var oldData []byte

	oldData, err = ioutil.ReadFile(filename)
	// read error occurred thus createFile should be called
	if err != nil {
		return
	}

	if bytes.Compare(oldData, newData) != 0 {
		err = ioutil.WriteFile(filename, newData, perm)
		changed = true
	}

	return
}

func deleteFile(filename string) (err error) {
	// this is just a vanity wrapper
	err = nil

	err = os.Remove(filename)

	return
}

func writeJournal() (err error){
	err = nil
	var buf bytes.Buffer

	for k := range journalMessages {
		buf.WriteString(k)
	}

	err = ioutil.WriteFile(journalFile, buf.Bytes(), os.FileMode(0666))

	journalMessages = make(map[string]bool)

	return
}

func addToJournal(message string) {
	journalMessages[message] = true
}

func main() {
	// Define a template.
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin hat", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for i, r := range recipients {
		filename := fmt.Sprintf("/tmp/%d.txt", i)
		var buf bytes.Buffer
		err := t.Execute(&buf, r)
		check(err)

		changed, err := updateFile(filename, buf.Bytes(), os.FileMode(0666))
		if changed {
			addToJournal(fmt.Sprintf("File: %s changed...\n", filename))

		}
		if err != nil {
			err := createFile(filename, buf.Bytes(), os.FileMode(0666))
			check(err)
			addToJournal(fmt.Sprintf("File: %s created...\n", filename))
		}
	}

	writeJournal()
}
