package named

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"net"

	"github.com/bitvector2/testgo/utils"
)

const (
	aclListTemplate = `
    {{- range .Acls }}
    {{ . }}
    {{- end }}
`
)

type AclList struct {
	template *template.Template
	filename string
	Acls     map[string]Acl
}

func NewAclList(outputDir string) *AclList {
	t := template.Must(template.New("aclListTemplate").Parse(aclListTemplate))

	return &AclList{
		template: t,
		filename: fmt.Sprintf("%s/named.conf.acllist", outputDir),
		Acls:     make(map[string]Acl, 0),
	}
}

func (al AclList) String() string {
	var buf bytes.Buffer
	err := al.template.Execute(&buf, al)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (al *AclList) Add(aclname string) {
	al.Acls[aclname] = *NewAcl(aclname)
}

func (al *AclList) Delete(aclname string) {
	delete(al.Acls, aclname)
}

func (al AclList) Save() bool {
	var buf bytes.Buffer
	buf.WriteString(al.String())
	dirty, err := utils.UpdateFile(al.filename, buf.Bytes(), os.FileMode(0666))
	utils.Check(err)
	return dirty
}

func (al AclList) Contains(aclname string) bool {
	for k, _ := range al.Acls {
		if k == aclname {
			return true
		}
	}
	return false
}

func (al AclList) Zombies() []string {
	aclnames := make([]string, 0)
	for k, v := range al.Acls {
		if v.Len() == 0 {
			aclnames = append(aclnames, k)
		}
	}
	return aclnames
}

func (al AclList) AddElement(key string, aclname string, ip string) {
	obj := al.Acls[aclname]
	obj.Add(key, *NewCidrAddress(net.ParseIP(ip), net.IPv4Mask(255, 255, 255, 255)))
}

func (al AclList) DelElement(key string) {
	for _, v := range al.Acls {
		if v.Contains(key) {
			v.Delete(key)
		}
	}
}
