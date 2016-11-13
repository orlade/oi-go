.PHONY : all build test watch stop

all: build test

build:
	go install
	go build -i -o cli/oi cli/main.go
	mv cli/oi $(GOPATH)/bin/oi

test:
	go test

watch:
	if [ "$(which watch)" == "" ]; then npm install -g watch; fi
	watch 'make build && make test' . --ignoreDotFiles > .watch.log 2>&1 &

stop:
	ps -ef | awk '/[n]ode.*watch/ {print $$2}' | xargs kill
