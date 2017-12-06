package main

import (
	"fmt"
	"net"

	"github.com/bitvector2/testgo/named"
	"github.com/bitvector2/testgo/utils"
)

const (
	version = "1.0.0"
)

func main() {
	//outputDir := flag.String("outputdir", "", "Absolute path to the output directory")
	//kubeConfig := flag.String("kubeconfig", "", "Absolute path to the Kubernetes config file")
	//masterURL := flag.String("masterurl", "", "URL to Kubernetes API server")
	//flag.Parse()
	//
	//var err error
	//var buf bytes.Buffer
	//buf.WriteString("// empty\n")
	//
	//err = utils.CreateFile(fmt.Sprintf("%s/named.conf.acllist", *outputDir), buf.Bytes(), os.FileMode(0666))
	//utils.Check(err)
	//
	//err = utils.CreateFile(fmt.Sprintf("%s/named.conf.viewlist", *outputDir), buf.Bytes(), os.FileMode(0666))
	//utils.Check(err)
	//
	//// Create our custom controller
	//c := pods.New(*kubeConfig, *masterURL)
	//
	//// Start our custom controller
	//stop := make(chan struct{})
	//defer close(stop)
	//go c.Run(1, stop)

	acl1 := named.NewAcl("acl1")
	for i := 1; i < 3; i++ {
		ip, subnet, err := net.ParseCIDR("192.168.8.1/24")
		utils.Check(err)
		acl1.AddElement(*named.NewCidrAddress(ip, subnet.Mask))
	}
	view1 := named.NewView("view1", *acl1)
	view1.AddZone(*named.NewZone("zone1", "/tmp/zone1.txt"))

	acl2 := named.NewAcl("acl2")
	for i := 1; i < 3; i++ {
		ip, subnet, err := net.ParseCIDR("192.168.9.1/24")
		utils.Check(err)
		acl2.AddElement(*named.NewCidrAddress(ip, subnet.Mask))
	}
	view2 := named.NewView("view2", *acl2)
	view2.AddZone(*named.NewZone("zone2", "/tmp/zone2.txt"))

	aclList := named.NewAclList("/tmp/acllist.txt")
	aclList.AddAcl(*acl1)
	aclList.AddAcl(*acl2)

	fmt.Print(aclList)

	viewList := named.NewViewList("/tmp/viewlist.txt")
	viewList.AddView(*view1)
	viewList.AddView(*view2)

	fmt.Print(viewList)

	// Run forever
	//log.Infoln("testgo version: " + version + " started...")
	//select {}

}
