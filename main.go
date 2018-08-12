package main

import (
	"fmt"
	"github.com/unquabain/gibberish/config"
	"github.com/unquabain/gibberish/lexicon"
	"github.com/unquabain/gibberish/server"
)

func main() {
	config.Init()	
	lexicon, err := lexicon.NewLexicon(".")
	if err != nil {
		fmt.Errorf("Could not create a new lexicon: %v", err)
		return
	}
	server := server.NewServer(lexicon)
	if err != nil {
		fmt.Errorf("Could not create the server: %v", err)
	} else {
		server.Serve(config.Port)
	}
}
