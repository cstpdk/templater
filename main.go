package main

import (
        "encoding/json"
        "fmt"
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
        Tmplfile *os.File
        Datafile *os.File
        Data     map[string]interface{} // Read data
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
func findTemplateData(tmplfile *os.File, datafile *os.File) map[string]interface{} {
        var data interface{}
        if datafile == nil {
                return make(map[string]interface{})
        }
        b, _ := ioutil.ReadFile(datafile.Name())
        json.Unmarshal(b, &data)
        m := data.(map[string]interface{})

        //for k, v := range m {
        //        switch vv := v.(type) {
        //        case string:
        //                fmt.Println("\n", k, "is string", vv)
        //        case int:
        //                fmt.Println("\n", k, "is int", vv)
        //        case float64:
        //                fmt.Println("\n", k, "is a double", vv)
        //        case float32:
        //                fmt.Println("\n", k, "is a double", vv)
        //        case []interface{}:
        //                fmt.Println("\n", k, "is an array:")
        //                for i, u := range vv {
        //                        fmt.Println(i, u)
        //                }
        //        default:
        //                fmt.Println(k, "is of a type I don't know how to handle")
        //                fmt.Printf("%v:%t", v, vv)
        //        }
        //}

        return m
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

// For all templates, create them in an appropiate location
func generateTemplates(destinationDir string, tmpls []Template) {
        templ := tmpls[0] // TODO loop
        fmt.Printf("%v", templ)
        tmplBytes, _ := ioutil.ReadFile(templ.Tmplfile.Name())
        tmpl := template.Must(template.New("NAME").Parse(string(tmplBytes[:])))
        fmt.Printf("%v", templ.Data)
        destination, err := os.OpenFile(path.Join(destinationDir, "joe"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
        if err != nil {
                panic(err)
        }
        err = tmpl.Execute(destination, templ.Data)
        if err != nil {
                panic(err)
        }
}

func init() {
        MakeConfig(&config, "config.json")

        tmpls := findTemplates(config.Templates_dir)
        generateTemplates("./generated", tmpls)
}

func main() {

}
