package core

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
)

var Box packr.Box

func init() {
	Box = packr.NewBox("../templates")
}

func pass(e error) bool {
	if e != nil {
		fmt.Printf("%+v", e)
		return false
	}
	return true
}

// ParseTemplates parse and load all template in a subpath
func ParseTemplates(destPath string) error {
	e := Box.Walk(func(path string, file packd.File) error {
		isTpl, e := filepath.Match("*.tpl", file.Name())
		if pass(e) && isTpl {
			t := template.New("test")
			t, e = t.Parse(file.String())
			if pass(e) {
				pass(t.Execute(os.Stdout, nil))
			}
		}
		return e
	})
	pass(e)
	return nil
}
