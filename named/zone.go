package named

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/bitvector2/testgo/utils"
)

const (
	zoneTemplate = `
    zone {{ .Name }} {
        type master;
        file {{ .Filename }};
    };
`
)

type Zone struct {
	Name            string
	template        *template.Template
	Filename        string
	ResourceRecords []ResourceRecord
}

func NewZone(name string, outputDir string) *Zone {
	t := template.Must(template.New("zoneTemplate").Parse(zoneTemplate))

	return &Zone{
		Name:            name,
		template:        t,
		Filename:        fmt.Sprintf("%s/db.%s", outputDir, name),
		ResourceRecords: make([]ResourceRecord, 0),
	}
}

func (z *Zone) String() string {
	var buf bytes.Buffer
	err := z.template.Execute(&buf, z)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (z *Zone) SprintResourceRecords() string {
	var buf bytes.Buffer
	for i := 0; i < len(z.ResourceRecords); i++ {
		buf.WriteString(z.ResourceRecords[i].String())
	}
	return buf.String()
}

func (z *Zone) AddResourceRecord(rr ResourceRecord) {
	z.ResourceRecords = append(z.ResourceRecords, rr)
}

func (z *Zone) Save() {
	var buf bytes.Buffer
	buf.WriteString(z.SprintResourceRecords())
	utils.CreateFile(z.Filename, buf.Bytes(), os.FileMode(0666))
}
