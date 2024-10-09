
all: run

install: deploy
	mkdir -p ~/.config/promptorium && cp -r conf/* ~/.config/promptorium && \
	mkdir -p ~/.config/promptorium/shell && cp shell/* ~/.config/promptorium/shell

deploy: build
	sudo cp build/promptorium-linux-amd64 /usr/local/bin/promptorium && sudo chmod +x /usr/local/bin/promptorium

run: build
	./build/promptorium-linux-amd64

build:
	cd src && \
	bash ../scripts/build.bash

build-deb: build
	bash scripts/deb/build-deb.bash
clean:
	if [ -d build ]; then rm -rf build; fi


.PHONY: build run clean
