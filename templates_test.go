package main

import (
        "testing"
)

func TestCapitalize(t *testing.T) {
        if Capitalize("john") != "John" {
                t.Fatal("Not good")
        }
}
