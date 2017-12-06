package named

import (
	"bytes"
	"text/template"
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

func NewZone(name string, filename string) *Zone {
	t := template.Must(template.New("zoneTemplate").Parse(zoneTemplate))

	return &Zone{
		Name:            name,
		template:        t,
		Filename:        filename,
		ResourceRecords: make([]ResourceRecord, 0),
	}
}

func (z Zone) String() string {
	var buf bytes.Buffer
	err := z.template.Execute(&buf, z)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
