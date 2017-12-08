package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/bitvector2/aced-dns/named"
	"github.com/bitvector2/aced-dns/pods"
	"github.com/bitvector2/aced-dns/utils"
	log "github.com/golang/glog"
)

const (
	version = "1.0.0"
)

func main() {
	outputDir := flag.String("outputdir", "", "Absolute path to the output directory")
	kubeConfig := flag.String("kubeconfig", "", "Absolute path to the Kubernetes config file")
	masterURL := flag.String("masterurl", "", "URL to Kubernetes API server")
	flag.Parse()

	var err error
	var buf bytes.Buffer

	err = utils.CreateFile(fmt.Sprintf("%s/named.conf.acllist", *outputDir), buf.Bytes(), os.FileMode(0666))
	utils.Check(err)

	err = utils.CreateFile(fmt.Sprintf("%s/named.conf.viewlist", *outputDir), buf.Bytes(), os.FileMode(0666))
	utils.Check(err)

	aclList := named.NewAclList(*outputDir)
	viewList := named.NewViewList(*outputDir)

	// Create our custom controller
	c := pods.New(*kubeConfig, *masterURL, *aclList, *viewList, *outputDir)

	// Start our custom controller
	stop := make(chan struct{})
	defer close(stop)
	go c.Run(1, stop)

	log.Infoln("aced-dns version: " + version + " started...")
	select {}
}
