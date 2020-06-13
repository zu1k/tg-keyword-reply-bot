NAME=bot
BINDIR=bin
GOBUILD=CGO_ENABLED=1 go build -ldflags '-w -s' -trimpath

PLATFORM_LIST = \
	darwin-amd64 \
	linux-amd64 \
	freebsd-amd64

WINDOWS_ARCH_LIST = \
	windows-amd64


linux-amd64:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(BINDIR)/$(NAME)-$@


clean:
	rm $(BINDIR)/*