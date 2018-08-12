package lexicon

import (
	"fmt"
	"math/rand"
	"regexp"
	"bytes"
	"strings"
	"github.com/unquabain/gibberish/config"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type WordList map[string]int

type Lexicon struct {
	List WordList `yaml:"words"`
	Path string   `yaml:"path"`
}

func NewLexicon(path string) (*Lexicon, error) {
	buff, err := config.Templates.MustBytes(filepath.Join(path, "words.yaml"))
	if err != nil {
		return nil, fmt.Errorf("Couldn't read input file: %v", err)
	}
	output := Lexicon{}
	
	var list WordList
	yaml.Unmarshal(buff, &list)
	output.List = list
	output.Path = path
	return &output, nil
}

func (this Lexicon) Choose() (winner string) {
	sum := 0
	for word, weight := range this.List {
		sum += weight
		n := rand.Intn(sum)
		if n < weight {
			winner = word
		}
	}
	return
}

func (this Lexicon) follow_path(path string) (string, error) {
	new_path := filepath.Clean(filepath.Join(this.Path, path))

	l, err := NewLexicon(new_path)
	if err != nil {
		return "", err
	}
	out, err := l.Evaluate()
	return out, err
}

func (this Lexicon) Evaluate() (string, error) {
	template := this.Choose()
	var output bytes.Buffer
	pattern, err := regexp.Compile("{{(.*?)}}")
	if err != nil {
		return "", fmt.Errorf("Could not compile pattern: %v", err)
	}
	end := 0
	for _, offsets := range pattern.FindAllStringSubmatchIndex(template, -1) {
		output.WriteString(template[end:offsets[0]])
		subpath := template[offsets[2]:offsets[3]]
		evalled, err := this.follow_path(strings.Trim(subpath, " "))
		if err != nil {
			return output.String(), fmt.Errorf("Couldn't evaluate path subpath: %v", err)
		}
		output.WriteString(evalled)
		end = offsets[1]
	}
	output.WriteString(template[end:])
	return output.String(), nil
}

