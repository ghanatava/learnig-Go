install: 
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: install
	swagger generate spec -o ./swagger.yml

.PHONY: swagger