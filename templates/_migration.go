package main

import (
        "github.com/eaigner/hood"
)

type {{.NameCapitalized}}s struct {
        Id       hood.Id `sql:"pk"`
        Name string  `validate:"presence"`
        Created  hood.Created
        Updated  hood.Updated
}

func (m *M) Create{{.NameCapitalized}}Table_{{.Timestamp}}_Up(hd *hood.Hood) {
        hd.CreateTable(&{{.NameCapitalized}}s{})
}

func (m *M) Create{{.NameCapitalized}}Table_{{.Timestamp}}_Down(hd *hood.Hood) {
        hd.DropTable(&{{.NameCapitalized}}s{})
}
