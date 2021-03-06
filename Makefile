.PHONY: build
build:
	go build genpasshash.go

.PHONY: deps
deps:
	go get "golang.org/x/crypto/ssh/terminal"
	go get "github.com/amoghe/go-crypt"

.PHONY: clean
clean:
	- rm genpasshash
	- rm *.exe
