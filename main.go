package main

import (
        //"fmt"
        "io/ioutil"
        "os"
        "path"
        "regexp"
)

var (
        config    Config
        templates []Template
)

// Template type, holds location of template, data etc.

type Template struct {
        Tmplfile *os.File
        Datafile *os.File
        Data     *interface{} // Read data
}

// Check whether the filename provided corresponds to being a template
// file
func isTemplateFile(filename string) bool {
        matched, err := regexp.MatchString(config.Template_regexp, filename)
        if err != nil {
                panic("The regexp for finding templates is wrong")
        }
        return matched
}

// Find all templates based on directory containing them
func findTemplateFiles(dirname string) []*os.File {
        files, err := ioutil.ReadDir(dirname)

        if err != nil {
                panic("Can't read template directory")
        }

        templateFileNames := []*os.File{}

        for _, e := range files {
                fName := e.Name()
                // If it conforms to the template layout
                if isTemplateFile(fName) {
                        file, _ := os.Open(path.Join(dirname, fName))
                        // Add it to the list
                        templateFileNames = append(templateFileNames,
                                file)
                }
        }
        return templateFileNames
}

// Check that the given template filename corresponds to the given data
// filename
func isDataForTemplate(tmplFilename string, dataFilename string) bool {
        matched, err := regexp.MatchString(
                tmplFilename+`\.json`, "_"+dataFilename)
        if err != nil {
                panic("The regexp for finding template data is wrong")
        }
        return matched
}

// Find datafile for a template
func findTemplateDataFile(dirname string, tmplFilename string) *os.File {
        files, err := ioutil.ReadDir(dirname)

        if err != nil {
                panic("Can't read template directory")
        }

        for _, e := range files {
                fName := e.Name()
                if isDataForTemplate(tmplFilename, fName) {
                        file, err := os.Open(path.Join(dirname, fName))
                        if err != nil {
                                panic(err)
                        }
                        return file
                }
        }

        return nil
}

// Extract correct data from a data file in the system
func findTemplateData(tmplfile *os.File, datafile *os.File) *interface{} {
        //print(tmplfile.Name(), " ", datafile.Name(), "\n")
        iface := new(interface{})
        return iface
}

// Fill templates with correct values for all templates
func findTemplates(dirname string) []Template {

        tmplFiles := findTemplateFiles(dirname)

        templates := []Template{}
        for _, e := range tmplFiles {
                datafile := findTemplateDataFile(dirname, e.Name())
                templates = append(templates, Template{
                        Tmplfile: e,
                        Datafile: datafile,
                        Data:     findTemplateData(e, datafile),
                })
        }

        return templates
}

// For all templates, create their file in generated

func init() {
        MakeConfig(&config, "config.json")

        //findTemplates(config.Templates_dir)
}

func main() {

}
