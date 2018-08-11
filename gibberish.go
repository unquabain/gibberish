package main

import (
	"fmt"
	"unquabain/gibberish/config"
	"unquabain/gibberish/lexicon"
)

func main() {
	config.Init()	
	lexicon, err := lexicon.NewLexicon(config.TemplateRoot)
	if err != nil {
		fmt.Errorf("Could not create a new lexicon: %v", err)
		return
	}
	ret, err := lexicon.Evaluate()
	if err != nil {
		fmt.Errorf("Could not evaluate: %v", err)
	} else {
		fmt.Println(ret)
	}
}
