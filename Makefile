PACKR=$(GOPATH)/bin/packr
GLIDE=/usr/local/bin/glide
SASSC=/usr/local/bin/sassc

run: gibberish
	./gibberish

gibberish: main.go lexicon/lexicon.go config/config.go server/server.go web/css/index.css web/js/index.js web/index.html $( shell find templates -type f ) $(PACKR) glide.lock
	packr build github.com/unquabain/gibberish

glide.lock: glide.yaml /usr/local/bin/glide
	$(GLIDE) install

web/css/%.css: scss/%.scss $(SASSC)
	$(SASSC) $<  > $@

$(GLIDE):
	brew install glide

$(PACKR):
	go get -u github.com/gobuffalo/packr/...

$(SASSC):
	brew install sassc

.PHONY: run
