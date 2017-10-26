export GOPATH=$(CURDIR)/.go

APP_NAME = voucher
DEBIAN_TMP = $(CURDIR)/deb
VERSION = `$(CURDIR)/out/$(APP_NAME) -v | cut -d ' ' -f 3`

$(CURDIR)/out/$(APP_NAME): $(CURDIR)/src/main.go
	go build -o $(CURDIR)/out/$(APP_NAME) $(CURDIR)/src/main.go

dep-install:
	go get github.com/golang/glog
	go get github.com/zenazn/goji
	go get github.com/zenazn/goji/graceful
	go get github.com/zenazn/goji/web
	go get github.com/gorilla/sessions
	go get github.com/dim13/unifi

fmt:
	gofmt -s=true -w $(CURDIR)/src

run:
	go run $(CURDIR)/src/main.go

clean:
	rm -f $(CURDIR)/out/*

clean-deb:
	rm -fr $(DEBIAN_TMP)
	rm -f $(CURDIR)/*.deb

debug:
	echo $(GOPATH)
