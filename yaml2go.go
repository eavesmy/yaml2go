package yaml2go

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"reflect"
	"regexp"
	"strings"
)

type Yaml2Struct struct {
	PkgName    string
	StructName string
	Data       []byte
	StructStr  string
}

// NewStruct creates a Yaml2Struct to process the convert
func NewStruct(pkgName, structName string, data []byte) *Yaml2Struct {
	return &Yaml2Struct{PkgName: pkgName, StructName: structName, Data: data, StructStr: ""}
}

// DoYaml2Struct convert yaml string to go struct
func (y *Yaml2Struct) DoYaml2Struct() (err error) {
	var obj map[string]interface{}
	err = yaml.Unmarshal(y.Data, &obj)
	if err != nil {
		return
	}
	y.StructStr = fmt.Sprintf("package %s\n\n", y.PkgName)
	y.StructStr = fmt.Sprintf("%stype %s struct {", y.StructStr, y.StructName)
	for k1, v1 := range obj {
		y.procKV(1, k1, v1)
	}
	y.StructStr = fmt.Sprintf("%s\n}", y.StructStr)
	return
}

func (y *Yaml2Struct) procMap(level int, k string, v interface{}) {
	y.StructStr = fmt.Sprintf("%s\n%s%s struct {", y.StructStr, strings.Repeat("\t", level), toCample(k))
	for k1, v1 := range v.(map[string]interface{}) {
		y.procKV(level+1, k1, v1)
	}
	y.StructStr = fmt.Sprintf("%s\n%s} `yaml:\"%s\"`", y.StructStr, strings.Repeat("\t", level), toCample(k))
}

func (y *Yaml2Struct) procKV(level int, k string, v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Map:
		y.procMap(level, k, v)
		break
	case reflect.Slice:
		y.StructStr = fmt.Sprintf("%s\n%s%s\t []string `yaml:\"%s\"`", y.StructStr, strings.Repeat("\t", level), toCample(k), k)
		break
	default:
		y.StructStr = fmt.Sprintf("%s\n%s%s\t %s `yaml:\"%s\"`", y.StructStr, strings.Repeat("\t", level), toCample(k), reflect.TypeOf(v), k)
	}
}

func toCample(str string) string {
	re := regexp.MustCompile(`[\d,a-zA-Z]*`)
	_strs := re.FindAllStringSubmatch(str, -1)
	var strs []string
	for _, si := range _strs {
		if len(si) > 0 && len(si[0]) > 0 {
			strs = append(strs, strings.ToUpper(string(si[0][0]))+si[0][1:])
		}
	}
	return strings.Join(strs, "")
}
