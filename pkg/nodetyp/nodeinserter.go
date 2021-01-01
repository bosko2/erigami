package nodetyp

import (
	"dev.azure.com/markobosnjak20/erigami/_git/erigami/configs/nodetypes"
	"dev.azure.com/markobosnjak20/erigami/_git/erigami/pkg/node"
	"gopkg.in/yaml.v2"
)

type YamlDat struct {
	Typname string        `yaml:"Typname"`
	Out     map[int][]int `yaml:"Out"`
}

func RunGlobal(data []YamlDat) {
	outs := make(map[int][]interface{})
	for i, d := range data {
		var tren node.Node = node.NewNode(nodetypes.Types[d.Typname], outs[i])
		inter := tren.Run()
		a := d.Out
		println(a)
		if interface{}(d.Out) != nil {
			for y, c := range d.Out {
				outs[y] = nil
				if c != nil {
					for z, p := range c {
						outs[y] = append(outs[y], nil)
						if interface{}(p) != nil {
							outs[y][z] = inter[p]
						}
					}
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
