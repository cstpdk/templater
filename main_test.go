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

        tmpl, _ := os.Open(path.Join(templatesDir, "_john.html"))
        data, _ := os.Open(path.Join(templatesDir, "john.html.json"))
        print(tmpl, data)
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
        print(tmpl, "\n", data, "\n")
        findTemplateData(tmpl, data)
}