package config

import (
	"math/rand"
	"time"
	"github.com/gobuffalo/packr"
)

var Templates packr.Box
var Web packr.Box

func Init () {
	Templates = packr.NewBox("../templates")
	Web = packr.NewBox("../web")

	rand.Seed(time.Now().UTC().UnixNano())
}
