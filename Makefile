# Makefile for Serverless Go

# Define the "phony" `make` commands. Phony `make` commands are commands that
# execute shell statements rather than transform files. It is good practice to
# let `make` know which commands are phony so that they work the same way
# (i.e., execute the shell statements defined in this Makefile) whether or not
# a file with the same name exists. For example, defining `clean` as phony
# causes `make clean` to work the same way regardless of whether a file named
# "clean" exists in the directory. For more information, see:
# https://stackoverflow.com/a/2145605
.PHONY: check-env \
	compile-and-deploy-lambda
	compile-lambda \
	deploy-lambda \
	deps \
	format \
	lint \

compile-and-deploy-lambda: compile-lambda deploy-lambda

compile-lambda:
	@rm -rf dist
	@mkdir -p dist
	@echo Compiling...
	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o dist/webservice ./cmd/webservice
	@echo Compiled for lambda
	@echo Zipping...
	@zip ./dist/webservice.zip ./dist/webservice

deploy-lambda:
	@echo Deploying to lambda...
	@aws lambda update-function-code --function-name jbcc-brc-api-go --zip-file fileb://./dist/webservice.zip

deps:
	GO111MODULE=on go mod vendor

format:
	go fmt ./...

lint:
	go vet ./...