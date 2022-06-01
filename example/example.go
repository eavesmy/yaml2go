package example

import (
	"github.com/fengzxu/yaml2go"
	"log"
	"os"
)

// GenConfigFile example for generate go struct from a config file (yaml)
// tips:you can use 'go:generate' to generate .go file for structs before starting logic codes.
func GenConfigFile() {
	var (
		configStructName = "Config"
		configFile       = "config.yaml" // more complicate yaml,try k3s.yaml
		goFile           = "config.go"
		goPackage        = "Config"
	)
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal("error on read YAML file:", err.Error())
	}
	y2s := yaml2go.NewStruct(goPackage, configStructName, data)
	err = y2s.DoYaml2Struct()
	if err != nil {
		log.Fatal("generate go file failed:", err.Error())
	}
	err = os.WriteFile(goFile, []byte(y2s.StructStr), 0644)
	if err != nil {
		log.Fatal("write go file failed:", err.Error())
	}
	log.Println("generate", goFile, "Done.")
}
