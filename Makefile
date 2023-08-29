.DEFAULT_GOAL := all

all: build_client start


## start: starts the server
start:
	@echo "Starting server"
	cd ./server && go run .

# build_client: builds the client as a linux executable with name calc
build_client:
	@echo "Building the client"	
	cd ./client && env GOOS=linux CGO_ENABLED=0 go build -o calc . && mv calc ../