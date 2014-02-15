package main

import (
        "io/ioutil"
        "os"
        "strings"
        "text/template"
)

type TemplateData struct {
        Name            string
        NameCapitalized string
        Timestamp       int64
}

func Capitalize(str string) string {
        return strings.ToUpper(string(str[0])) + str[1:]
}

func createTemplate(
        destination *os.File, tmplData TemplateData, tmplFile string) {

        tmplBytes, _ := ioutil.ReadFile(tmplFile)

        tmpl := template.Must(
                template.New("aname").Parse(string(tmplBytes[:])))

        tmpl.Execute(destination, tmplData)
}

//func createMigrationTemplate(
//        destination *os.File, tmplData TemplateData) {
//
//        createTemplate(
//                destination, tmplData, "cmd/templates/_migration.go")
//}
