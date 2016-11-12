build:
	go build -i -o oi cli/main.go
	mv oi $(GOPATH)/bin/oi

watch:
	npm install -g watch
	watch 'make build' . -do > .watch.log 2>&1 &

stop:
	ps -ef | awk '/[n]ode.*watch/ {print $$2}' | xargs kill
