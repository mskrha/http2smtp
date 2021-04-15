BUILD		= http2smtp
VERSION		?= 0.0
PKG		?= $(BUILD)_$(VERSION)_amd64
OUTPUT		= $(BUILD)
REGISTRY	= 'registry.docker.srv.sobotiste.info'

all: clean format build

clean:
	rm -f $(OUTPUT)

format:
	go fmt

build:
	CGO_ENABLED=0 go build -ldflags "-X main.version=$(VERSION)" -o $(OUTPUT) *.go

temp:
	$(eval TMPDIR = $(shell mktemp -d))
	$(eval OUTPUT = $(TMPDIR)/usr/sbin/$(BUILD))
	mkdir -p $(TMPDIR)/etc/systemd/system
	mkdir -p $(TMPDIR)/etc/http2smtp
	mkdir -p $(TMPDIR)/usr/sbin

files: temp
	cp -av ../DEBIAN $(TMPDIR)/
	cp -av ../config.json $(TMPDIR)/etc/http2smtp/
	cp -av ../http2smtp.service $(TMPDIR)/etc/systemd/system/

deb: clean format files build
	$(eval SIZE = $(shell du -k $(OUTPUT) | sed 's/^\([0-9]\+\).*/\1/'))
	$(shell sed -i "s/INSTALL_SIZE/$(SIZE)/" $(TMPDIR)/DEBIAN/control)
	$(shell sed -i "s/INSTALL_VERSION/$(VERSION)/" $(TMPDIR)/DEBIAN/control)
	dpkg-deb --build $(TMPDIR) $(PKG).deb
	rm -rf $(TMPDIR)

docker: clean format build
	docker build --compress --tag $(REGISTRY)/$(BUILD):$(VERSION) --tag $(REGISTRY)/$(BUILD):latest .
	docker push $(REGISTRY)/$(BUILD):$(VERSION)
	docker push $(REGISTRY)/$(BUILD):latest
	rm -f $(BUILD)
