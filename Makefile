all: bin/templater

.PHONY: docs

bin/templater: *.go templates/
	go build -o bin/templater github.com/cstpdk/templater

docs: docs/*.md
	find docs/ -name "*.md" -exec pandoc -o {}.pdf {} \;

clean:
	rm docs/*.pdf
	rm bin/templater
