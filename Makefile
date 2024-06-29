
prepare:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest;\
	go install github.com/automation-co/husky@latest;\
	(husky init || true) && husky install;