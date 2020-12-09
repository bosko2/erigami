package node

import (
	"dev.azure.com/markobosnjak20/erigami/_git/erigami/configs/nodetypes"
	"gopkg.in/yaml.v2"
)

type YamlDat struct {
	typname string
	out     map[int][]int
}

func RunGlobal(data []YamlDat) {
	outs := make(map[int][]interface{})
	for i, d := range data {
		var tren Node = NewNode(nodetypes.Types[d.typname], outs[i])
		inter := tren.Run()
		for y, c := range d.out {
			for z, d := range c {
				if d != int(nil) {
					outs[y][z] = inter[z]
				}
			}
		}
	}
}

func getFromYaml(dat []byte) (v []YamlDat) {
	err := yaml.Unmarshal(dat, &v)
	if err != nil {
		panic(err)
	}
	return
}
