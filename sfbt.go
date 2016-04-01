package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/tdewolff/minify/css"

	"github.com/tdewolff/minify"
)

var cssFolder string
var cssTargetFile string
var cssFile *os.File
var cssWriter io.WriteCloser

func main() {
	config, err := toml.LoadFile("conf.toml")
	check(err)
	if config.Has("sfbt.folder") {
		cssFolder = config.Get("sfbt.folder").(string)
	}
	if config.Has("sfbt.targetFile") {
		cssTargetFile = config.Get("sfbt.targetFile").(string)
	}

	min := minify.New()
	min.AddFunc("text/css", css.Minify)

	cssFile, err = os.Create(cssTargetFile)
	check(err)

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
