
all: run

install: build-deb
	sudo dpkg -i build/deb/promptorium_$(PROMPTORIUM_VERSION)-1_amd64.deb

build-deb: build
	bash scripts/build-deb.bash

build: clean
	mkdir -p build
	cd src && \
	bash ../scripts/build.bash

clean:
	if [ -d build ]; then rm -rf build; fi


.PHONY: build run clean
