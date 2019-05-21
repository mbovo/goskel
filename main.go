package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

var (
	templatePath = "./templates"
)

func pass(e error) bool {
	if e != nil {
		fmt.Printf("%+v", e)
		return false
	}
	return true
}

// ParseTemplate parse and load all template in a subpath
func ParseTemplate(rootPath string) error {

	type voidData struct {
		M bool
	}
	data := voidData{false}

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if pass(err) {
			isTemplate, err := filepath.Match("*.tpl", info.Name())
			if pass(err) {
				if !info.IsDir() && isTemplate {
					t := template.New("test")
					t, err := template.ParseFiles(path)
					if pass(err) {
						t.Execute(os.Stdout, data)
					}
				}
			}
		}
		return err
	})
	return err
}

// Hallo is my hallo
func Hallo() string {
	return "Hello, world."
}

func main() {
	fmt.Printf("%s \n", Hallo())
	e := ParseTemplate("./templates")
	pass(e)
}
