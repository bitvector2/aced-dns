package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/bitvector2/testgo/pods"
	log "github.com/golang/glog"
)

const (
	version = "1.0.0"

	letter = `
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
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createFile(filename string, newData []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, newData, perm)
}

func updateFile(filename string, newData []byte, perm os.FileMode) (bool, error) {
	var err error
	var oldData []byte

	// read error occurred thus createFile should be called
	oldData, err = ioutil.ReadFile(filename)
	if err != nil {
		return false, err
	}

	if bytes.Compare(oldData, newData) != 0 {
		err = ioutil.WriteFile(filename, newData, perm)
		return true, err
	}

	return false, nil
}

func deleteFile(filename string) error {
	return os.Remove(filename)
}

func main() {
	outputDir := flag.String("outputdir", "", "Absolute path to the output directory")
	kubeConfig := flag.String("kubeconfig", "", "Absolute path to the Kubernetes config file")
	masterURL := flag.String("masterurl", "", "URL to Kubernetes API server")
	flag.Parse()

	var err error
	var buf bytes.Buffer
	buf.WriteString("// empty\n")

	err = createFile(fmt.Sprintf("%s/named.conf.acllist", *outputDir), buf.Bytes(), os.FileMode(0666))
	check(err)

	err = createFile(fmt.Sprintf("%s/named.conf.viewlist", *outputDir), buf.Bytes(), os.FileMode(0666))
	check(err)

	// Create our custom controller
	c := pods.New(*kubeConfig, *masterURL)

	// Start our custom controller
	stop := make(chan struct{})
	defer close(stop)
	go c.Run(1, stop)

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
		filename := fmt.Sprintf("%s/%d.txt", *outputDir, i)
		var buf bytes.Buffer
		err := t.Execute(&buf, r)
		check(err)

		changed, err := updateFile(filename, buf.Bytes(), os.FileMode(0666))
		if changed {
			log.Infof("File: %s changed...\n", filename)

		}
		if err != nil {
			err := createFile(filename, buf.Bytes(), os.FileMode(0666))
			check(err)
			log.Infof("File: %s created...\n", filename)
		}
	}

	// Run forever
	log.Infoln("testgo version: " + version + " started...")
	select {}

}
