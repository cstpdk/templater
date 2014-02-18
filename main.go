package main

import (
        "encoding/json"
        "io/ioutil"
        "os"
        "path"
        "regexp"
        "text/template"
)

var (
        config    Config
        templates []Template
)

// Template type, holds location of template, data etc.

type Template struct {
        Name            string
        Destination_dir string
        Tmplfile        *os.File
        Datafile        *os.File
        Data            map[string]interface{} // Read data
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
                path.Base(tmplFilename)+`\.json`, "_"+dataFilename)
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
func findTemplateData(tmplfile *os.File, datafile *os.File) map[string]interface{} {
        var data interface{}
        if datafile == nil {
                return make(map[string]interface{})
        }
        b, _ := ioutil.ReadFile(datafile.Name())
        json.Unmarshal(b, &data)
        m := data.(map[string]interface{})

        return m
}

// Create a Template struct based on values given in the datafile.
// Account for special/reserved values
func generateTemplateFromFiles(tmplfile *os.File,
        datafile *os.File) Template {

        // Data file for the template
        data := findTemplateData(tmplfile, datafile)

        // If data file specifies a name use that, otherwise use
        // file location as name
        var name string
        if n, present := data["Name"]; present {
                name = n.(string)
        } else {
                name = tmplfile.Name()
        }

        // If the data file specifies a location for the file use that,
        // Otherwise use a generic location
        var destination string
        if d, present := data["Destination"]; present {
                destination = d.(string)
        } else {
                destination = config.Default_output_dir
        }
        return Template{
                Destination_dir: destination,
                Name:            name,
                Tmplfile:        tmplfile,
                Datafile:        datafile,
                Data:            data,
        }
}

// Fill templates with correct values for all templates
func findTemplates(dirname string) []Template {

        tmplFiles := findTemplateFiles(dirname)

        templates := []Template{}
        for _, e := range tmplFiles {
                datafile := findTemplateDataFile(dirname, e.Name())
                tmpl := generateTemplateFromFiles(e, datafile)
                templates = append(templates, tmpl)
        }

        return templates
}

// Extract a go template and the file to put it in from a Template{}
func generateTemplate(templ Template) (*os.File,
        *template.Template, error) {

        tmplBytes, _ := ioutil.ReadFile(templ.Tmplfile.Name())
        tmpl := template.Must(template.New(templ.Name).Parse(string(tmplBytes[:])))
        destination := path.Join(templ.Destination_dir, templ.Name)
        destinationFile, err := os.OpenFile(destination, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

        if err != nil {
                return nil, nil, err
        }
        return destinationFile, tmpl, nil
}

// For all templates, create them in an appropiate location
func generateTemplates(tmpls []Template) {

        for _, e := range tmpls {
                file, tmpl, err := generateTemplate(e)

                if err != nil {
                        panic(err)
                }

                tmpl.Execute(file, e.Data)
        }

}

func init() {
        MakeConfig(&config, "config.json")
}

func main() {
        tmpls := findTemplates(config.Templates_dir)
        generateTemplates(tmpls)
}
