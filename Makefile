SHELL := /bin/bash

.PHONY: build
build:
	go build -o bin/pokedexcli cmd/pokedexcli/main.go

.PHONY: build_and_run
build_and_run: build
	./bin/pokedexcli

.PHONY: clean
clean:
	rm bin/*
