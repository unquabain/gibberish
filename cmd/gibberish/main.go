package main

import (
	"fmt"
	"unquabain/gibberish/config"
	"unquabain/gibberish/lexicon"
	"unquabain/gibberish/server"
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
		server.Serve(8181)
	}
}
