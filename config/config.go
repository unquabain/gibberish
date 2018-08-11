package config

import (
	"fmt"
	"flag"
	"math/rand"
	"time"
	"path/filepath"
)

var TemplateRoot string
var AbsRoot string

func Init () {
	var err error
	AbsRoot, err = filepath.Abs("./templates")
	if err != nil {
		fmt.Println("Couldn't find template path", err)
		return
	}

	flag.StringVar(&TemplateRoot, "template", AbsRoot, "A directory of lexicon files")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
}
