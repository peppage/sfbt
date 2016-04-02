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

	"github.com/tdewolff/minify"
)

var cssTargetFile string
var cssWriter io.WriteCloser

func main() {
	config, err := toml.LoadFile("conf.toml")
	if err != nil {
		log.Fatal("Missing conf.toml")
	}

	var cssFolder string

	if !config.Has("sfbt.cssFolder") {
		log.Fatal("config missing css folder")
	}

	if !config.Has("sfbt.targetCssFile") {
		log.Fatal("config missing target css file")
	}

	cssFolder = config.Get("sfbt.cssFolder").(string)
	cssTargetFile = config.Get("sfbt.targetCssFile").(string)

	min := minify.New()
	min.AddFunc("text/css", css.Minify)

	cssFile, err := os.Create(cssTargetFile)
	if err != nil {
		log.Fatal("Failed to create target css file")
	}

	cssWriter = min.Writer("text/css", cssFile)

	defer cssFile.Close()
	defer cssWriter.Close()

	filepath.Walk(cssFolder, handleCSS)
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
