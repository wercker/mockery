package tracery

import (
	"io/ioutil"
	"strings"
	"text/template"
)

type TemplateData struct {
	PkgName         string
	WrappedName     string
	WrappedFullName string // includes package
	WrapperName     string
	Methods         []TemplateFunction
	Imports         []TemplatePackage
}

type TemplatePackage struct {
	Name string
	Path string
}

type TemplateFunction struct {
	WrappedName  string
	WrapperName  string
	FunctionName string
	SigParams    string
	SigReturn    string
	Params       string
	CallParams   string
	Return       string
}

type GetObjInput struct {
	Foo bool
}

type GetObjOutput struct {
	Bar bool
}

type Traceme interface {
	GetObj(GetObjInput) GetObjOutput
}

var importTmpl string = `{{range .Imports}}import {{.Name}} "{{.Path}}"
{{end}}`

var replacements [][]string = [][]string{
	[]string{"//{{", "{{"},
	[]string{"package main", "package {{.PkgName}}"},
	[]string{`import "github.com/wercker/tracery/tracery"`, importTmpl},
	[]string{"TracingTraceme", "{{.WrapperName}}"},
	[]string{"tracery.Traceme", "{{.WrappedFullName}}"},
	[]string{"Traceme", "{{.WrappedName}}"},
	[]string{"GetObj(", "{{.FunctionName}}("},
	[]string{"methodInput tracery.GetObjInput", "{{.SigParams}}"},
	[]string{"tracery.GetObjOutput", "{{.SigReturn}}"},
	[]string{"(methodInput)", "({{.CallParams}})"},
	[]string{"methodInput", "{{.Params}}"},
	[]string{"methodOutput := ", "{{with .Return}}{{.}} := {{end}}"},
	[]string{"return methodOutput", "{{with .Return}}return {{.}}{{end}}"},
	[]string{"methodOutput", "{{.Return}}"},

	[]string{"blueprint/templates/service", "{{lower .Name}}"},
	[]string{"Blueprint", "{{title .Name}}"},
	[]string{"blueprint", "{{lower .Name}}"},
	[]string{"666", "{{.Port}}"},
	[]string{"667", "{{.Gateway}}"},
	[]string{"TiVo for VRML", "{{.Description}}"},
	[]string{"1996", "{{.Year}}"},
}

func replaceSentinels(s string) string {
	for _, x := range replacements {
		search, replace := x[0], x[1]
		s = strings.Replace(s, search, replace, -1)
	}
	return s
}

func GetTemplate(templatePath string) (*template.Template, error) {
	content, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	contentString := string(content)
	contentString = replaceSentinels(contentString)

	tmpl, err := template.New(templatePath).Funcs(Funcs).Parse(contentString)
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

func studly(s string) string {
	parts := strings.Split(s, "-")
	newParts := []string{}
	for _, part := range parts {
		newParts = append(newParts, strings.Title(part))
	}
	return strings.Join(newParts, "")
}

var Funcs template.FuncMap = template.FuncMap{
	// "package": func(input string) string { return strings.ToLower(input) },
	// "method":  func(input string) string { return strings.Title(input) },
	// "class":   func(input string) string { return strings.Title(input) },
	// "file":    func(input string) string { return strings.ToLower(input) },
	"title": studly,
	"lower": func(input string) string { return strings.ToLower(input) },
}
