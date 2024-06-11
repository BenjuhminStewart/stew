
OUT := stew
PKG := github.com/BenjuhminStewart/stew
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done
