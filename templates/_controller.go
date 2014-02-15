package api

import (
        //"bitbucket.org/leanalpha/qunizserver/models"
        "github.com/codegangsta/martini"
        "github.com/eaigner/hood"
)

// Performs routing related to {{.Name}}s. 
// The meat is in the respective functions in this file
func {{.NameCapitalized}}Route(r martini.Router) {

        r.Get("/{{.Name}}s", {{.NameCapitalized}}Index)

        r.Get("/{{.Name}}s/:id", {{.NameCapitalized}}Show)

        r.Put("/{{.Name}}s/:id", {{.NameCapitalized}}Update)
        r.Patch("/{{.Name}}s/:id", {{.NameCapitalized}}Update)

        r.Post("/{{.Name}}s", {{.NameCapitalized}}Create)

        r.Delete("/{{.Name}}s/:id", {{.NameCapitalized}}Destroy)
}

// GET /{{.Name}}s
//
// List all {{.Name}}s
//
// Returns
// Except if
func {{.NameCapitalized}}Index (db *hood.Hood) (int,string){
        return 501, "Not implemented"
}

// GET /{{.Name}}s/:id
//
// Show one {{.Name}}s
//
// Returns
// Except if
func {{.NameCapitalized}}Show (db *hood.Hood, params martini.Params) (int,string){
        //id := params["id"]
        return 501, "Not implemented"
}

// PUT /{{.Name}}s/:id
// PATCH /{{.Name}}s/:id
//
// Update {{.Name}} with id :id
//
// Returns
// Except if
func {{.NameCapitalized}}Update (db *hood.Hood, params martini.Params) (int,string){
        //id := params["id"]
        return 501, "Not implemented"
}

// POST /{{.Name}}s
//
// Create a new {{.Name}}
//
// Returns
// Except if
func {{.NameCapitalized}}Create (db *hood.Hood, params martini.Params) (int,string){
        return 501, "Not implemented"
}


// DELETE /{{.Name}}s
//
// Delete an exciting {{.Name}}
//
// Returns
// Except if
func {{.NameCapitalized}}Destroy (db *hood.Hood, params martini.Params) (int,string){
        //id := params["id"]
        return 501, "Not implemented"
}
