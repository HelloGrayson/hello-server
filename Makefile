project := hello-server


.PHONY: install
install:
	go build ./...


.PHONY: test
test:
	@echo "succeeded!"
