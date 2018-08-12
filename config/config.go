package config

import (
	"math/rand"
	"time"
	"github.com/gobuffalo/packr"
	"flag"
)

var Templates packr.Box
var Web packr.Box
var Port int

func Init () {
	Templates = packr.NewBox("../templates")
	Web = packr.NewBox("../web")

	flag.IntVar(&Port, "port", 8181, "The port to serve on")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
}
