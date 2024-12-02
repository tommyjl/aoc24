.PHONY: run
run:
	@go run .

.PHONY: watch
watch:
	find . -name "*.go" | entr -cr $(MAKE)
