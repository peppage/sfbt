# sfbt
A very simple static file build tool

sfbt is a build tool helper. It's best used in a make file to move and minify your static files. I didn't want to
install node for this one little task. It's not an exact solution but close enough for me.

## Install

Either

```
go get github.com/peppage/sfbt
```

or download executable and move into yuour $GOPATH/bin

## Usage
For most of my projects I use this example https://gist.github.com/Stratus3D/a5be23866810735d7413

Just have to add it to your build rule

```
build: vet
	gorazor ./tmpl ./tmpl
	sfbt
	go build -v
```

You also need to have a file "conf.toml" in the directory with the makefile and with these variables. You can have a path
in the targetCssFile so it writes the file where you want but the folder must be created.

```
[sfbt]
cssFolder = "css"
targetCssFile = "static/site.css"
jsFolder = "js"
targetJsFile = "static/site.js"
```
