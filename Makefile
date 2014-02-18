all: bin/templater

.PHONY: docs

bin/templater: *.go templates/
	go build -o bin/templater github.com/cstpdk/templater

clean:
	rm bin/templater
