package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/negrel/configue"
)

//go:embed gen.tpl
var goTemplate string

type Params struct {
	Package        string
	OptionArticle  string
	OptionName     string
	MethodReceiver string
	MethodType     string
	Out            string
}

func main() {
	params := Params{}

	figue := configue.New(
		"",
		configue.ContinueOnError,
		configue.NewEnv("CODEGEN"),
		configue.NewFlag(),
	)
	figue.StringVar(&params.Package, "package", "",
		"go package name of generated code")
	figue.StringVar(&params.OptionArticle, "option.article", "",
		"article to user before option name")
	figue.StringVar(&params.OptionName, "option.name", "",
		"name of option in go doc comments (e.g. flag, environmet variable)")
	figue.StringVar(&params.MethodReceiver, "method.receiver", "",
		"method receiver of generated function")
	figue.StringVar(&params.MethodType, "method.type", "",
		"method type of generated function")
	figue.StringVar(&params.Out, "out", "",
		"output file of generated code")

	err := figue.Parse()
	if err != nil {
		if err == flag.ErrHelp {
			return
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if params.OptionArticle == "" || params.OptionName == "" ||
		params.MethodReceiver == "" || params.MethodType == "" ||
		params.Out == "" {
		fmt.Fprintf(os.Stderr, "some options are not defined: %+v\n", params)
		os.Exit(1)
	}

	w, err := os.OpenFile(params.Out, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		println("failed to open output file:", err.Error())
		os.Exit(1)
	}

	tpl := template.Must(template.New("codegen").
		Funcs(sprig.FuncMap()).
		Parse(goTemplate))
	err = tpl.Execute(w, params)
	if err != nil {
		println("failed to execute template:", err.Error())
		os.Exit(1)
	}
}
