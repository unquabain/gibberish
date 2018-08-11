package config

import (
	"math/rand"
	"time"
	"github.com/gobuffalo/packr"
)

var Templates packr.Box

func Init () {
	Templates = packr.NewBox("../templates")

	rand.Seed(time.Now().UTC().UnixNano())
}
