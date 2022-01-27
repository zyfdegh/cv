GO=go

all: vend build test cover

vend:
	${GO} mod tidy && ${GO} mod vendor

build:
	${GO} build -mod=vendor

test:
	${GO} test -mod=vendor -timeout=90s -race ./... -coverprofile cover.out

cover:
	${GO} tool cover -func cover.out
