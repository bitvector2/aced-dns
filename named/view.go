package named

import (
	"bytes"
	"text/template"
)

const (
	viewTemplate = `
view {{ .Name }} {
    match-clients {
        {{ .Clients.Name }};
    };
    {{ range .Zones }}
    {{ . }}
    {{ end }}
};
`
)

type View struct {
	Name     string
	template *template.Template
	Clients  Acl
	Zones    []Zone
}

func NewView(name string, clientsAcl Acl) *View {
	t := template.Must(template.New("viewTemplate").Parse(viewTemplate))

	return &View{
		Name:     name,
		template: t,
		Clients:  clientsAcl,
		Zones:    make([]Zone, 0),
	}
}

func (v *View) String() string {
	var buf bytes.Buffer
	err := v.template.Execute(&buf, v)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (v *View) AddZone(zone Zone) {
	v.Zones = append(v.Zones, zone)
}

func (v *View) Save() {
	for z := 0; z < len(v.Zones); z++ {
		v.Zones[z].Save()
	}
}
