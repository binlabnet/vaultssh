VERSION=v0.1.4
TAGDESC="Fifth release (go mod support)"
BUILDTIME?=$$(date +%m-%d-%Y-%H:%M)
VERSIONSTRING=${VERSION}-${BUILDTIME}
GOFMT_FILES?=$$(find . -name '*.go')
export GO111MODULE=on

default: bin

all: fmt bin test

bin:
	go install -ldflags "-X github.com/richard-mauri/vaultssh/vs.VersionString=${VERSIONSTRING}"

test:
	go test github.com/richard-mauri/vaultssh/vs

fmt:
	gofmt -w $(GOFMT_FILES)

release:
	git tag -a ${VERSION} -m ${TAGDESC}
	RELVERSION=${VERSIONSTRING} goreleaser 

.PHONY: all bin default test fmt release
