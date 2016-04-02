package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/js"

	"github.com/tdewolff/minify"
)

var cssTargetFile string
var jsTargetFile string
var cssWriter io.WriteCloser
var jsWriter io.WriteCloser

func main() {
	config, err := toml.LoadFile("conf.toml")
	if err != nil {
		log.Fatal("Missing conf.toml")
	}

	if !config.Has("sfbt.cssFolder") {
		log.Fatal("config missing css folder")
	}

	if !config.Has("sfbt.targetCssFile") {
		log.Fatal("config missing target css file")
	}

	if !config.Has("sfbt.jsFolder") {
		log.Fatal("config missing js folder")
	}

	if !config.Has("sfbt.targetJsFile") {
		log.Fatal("config missing target js file")
	}

	cssFolder := config.Get("sfbt.cssFolder").(string)
	cssTargetFile = config.Get("sfbt.targetCssFile").(string)
	jsFolder := config.Get("sfbt.jsFolder").(string)
	jsTargetFile = config.Get("sfbt.targetJsFile").(string)

	min := minify.New()
	min.AddFunc("text/css", css.Minify)
	min.AddFunc("text/js", js.Minify)

	cssFile, err := os.Create(cssTargetFile)
	if err != nil {
		log.Fatal("Failed to create target css file")
	}

	jsFile, err := os.Create(jsTargetFile)
	if err != nil {
		log.Fatal("Failed to create target js file")
	}

	cssWriter = min.Writer("text/css", cssFile)
	jsWriter = min.Writer("text/js", jsFile)

	defer cssFile.Close()
	defer cssWriter.Close()
	defer jsFile.Close()
	defer jsWriter.Close()

	filepath.Walk(cssFolder, handleCSS)
	filepath.Walk(jsFolder, handleJs)
}

func handleCSS(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && strings.Contains(f.Name(), "css") {
		b, err := ioutil.ReadFile(path)
		check(err)

		_, err = cssWriter.Write(b)
		check(err)

	}
	return nil
}

func handleJs(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && strings.Contains(f.Name(), "js") {
		b, err := ioutil.ReadFile(path)
		check(err)

		_, err = jsWriter.Write(b)
		check(err)
	}
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
