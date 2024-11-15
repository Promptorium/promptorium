
all: run
	
install:
	sudo cp -f build/promptorium_$(PROMPTORIUM_VERSION)_linux_amd64 /usr/local/bin/promptorium
	cp -r conf/* ~/.config/promptorium
	cp -r shell/* ~/.config/promptorium/shell

install-deb: build-deb
	sudo apt --reinstall install ./build/deb/promptorium_$(PROMPTORIUM_VERSION)-1_amd64.deb

build-deb: build
	bash scripts/build-deb.bash

build: clean
	mkdir -p build
	cd src && \
	bash ../scripts/build.bash

clean:
	if [ -d build ]; then rm -rf build; fi


.PHONY: build run clean
