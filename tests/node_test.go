package tests

import (
	"dev.azure.com/markobosnjak20/erigami/_git/erigami/pkg/nodetyp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestNode(t *testing.T) {
	if _, err := os.Stat("./testdata.yaml"); os.IsNotExist(err) {
		log.Fatal("DOES NOT EGSIST")
	}
	data, err := ioutil.ReadFile("./testdata.yaml")
	if err != nil {
		panic(err)
	}
	var out []nodetyp.YamlDat
	err = yaml.Unmarshal(data, &out)
	if err != nil {
		panic(err)
	}
	nodetyp.RunGlobal(out)
}
