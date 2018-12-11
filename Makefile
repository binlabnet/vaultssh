VERSION=v0.1.1
TAGDESC="Second release"
BUILDTIME?=$$(date +%m-%d-%Y-%H:%M)
VERSIONSTRING=${VERSION}-${BUILDTIME}
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: bin

all: depends fmt bin test

depends:
	dep ensure -update

bin:
	go install -ldflags "-X github.com/richard-mauri/vaultssh/vs.VersionString=${VERSIONSTRING}"

test:
	go test github.com/richard-mauri/vaultssh/vs

fmt:
	gofmt -w $(GOFMT_FILES)

release:
	git tag -a ${VERSION} -m "${TAGDESC}"
#	RELVERSION=${VERSIONSTRING} goreleaser 

.PHONY: all bin default test fmt depends release
