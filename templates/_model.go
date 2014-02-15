package models

import (
        d "bitbucket.org/leanalpha/qunizserver/db"
)

type {{.NameCapitalized}}s d.{{.NameCapitalized}}s

func (entity {{.NameCapitalized}}s) John() {
        //return entity.Whatever
}
