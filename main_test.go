package main

import (
        "os"
        "path"
        "testing"
)

var (
        templatesCount int
        templatesDir   string
        tmpl           *os.File
        data           *os.File
)

func init() {
        templatesCount = 2
        templatesDir = "test_data/templates"

        tmpl, _ = os.Open(path.Join(templatesDir, "_john.html"))
        data, _ = os.Open(path.Join(templatesDir, "john.html.json"))
}

func TestIsTemplateFile(t *testing.T) {
        if !isTemplateFile("_name.html") {
                t.Fatal()
        }

        if isTemplateFile("iam_not_a_templatE_file.html") {
                t.Fatal()
        }
}
func TestFindTemplateFiles(t *testing.T) {
        tmpls := findTemplateFiles(templatesDir)
        // It should find only the one template "_john.html"
        if len(tmpls) != templatesCount {
                t.Fatalf("Found wrong amount of templates %d "+
                        "instead of %d", len(tmpls), templatesCount)
        }

        //        if tmpls[0].Name() != path.Join(templatesDir, "_john.html") {
        //                t.Fatalf("Found template was wrong, found %s", tmpls[0])
        //        }
}

func TestIsDataForTemplate(t *testing.T) {
        if !isDataForTemplate("_john.html", "john.html.json") {
                t.Fatal()
        }
        if isDataForTemplate("_john.html", "john.htmlgjson") {
                t.Fatal()
        }

        if isDataForTemplate("_sometemplate.html", "john.html.json") {
                t.Fatal()
        }
}

func TestFindTemplateDataFile(t *testing.T) {
        // It should find only the one template file for _john.html
        // and no data for _dataless_template.html

        if f := findTemplateDataFile(templatesDir,
                "_john.html"); f.Name() != path.Join(templatesDir,
                "john.html.json") {
                t.Fatal()
        }
        if f := findTemplateDataFile(templatesDir,
                "_dataless_template.html"); f != nil {
                t.Fatal()
        }
}

func TestFindTemplateData(t *testing.T) {
        res := findTemplateData(tmpl, data)

        if res["A_STRING"] != "string" {
                t.Fatal()
        }
        if res["lowercase_key"] != "ok" {
                t.Fatal()
        }
        if _, present := res["bogus_key"]; present {
                t.Fatal()
        }
}

func TestGenerateTemplates(t *testing.T) {

        tmpls := []Template{
                Template{
                        Name:     "joe",
                        Tmplfile: tmpl,
                        Datafile: data,
                        Data: map[string]interface{}{
                                "A_STRING": "joe",
                        },
                },
        }

        locationDir := "test_data/generated"
        filename := path.Join(locationDir, tmpls[0].Name)

        if err := os.Remove(filename); !os.IsNotExist(err) &&
                err != nil {

                t.Fatal("Couldn't delete previous file"+
                        "for some reason", err)
        }

        generateTemplates(locationDir, tmpls)

        if _, err := os.Stat(filename); os.IsNotExist(err) {
                t.Fatal("Template file was not generated")
        }
}
