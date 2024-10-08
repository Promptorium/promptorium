
all: run

install: deploy
	mkdir -p ~/.config/promptorium && cp -r conf/* ~/.config/promptorium && \
	mkdir -p ~/.config/promptorium/shell && cp shell/promptorium.bash ~/.config/promptorium/shell/promptorium.bash &&  \
	cp shell/promptorium.zsh ~/.config/promptorium/shell/promptorium.zsh

deploy: build
	sudo cp build/promptorium /usr/local/bin && sudo chmod +x /usr/local/bin/promptorium

run: build
	./build/promptorium

build:
	cd src && go build -o ../build/promptorium

clean:
	if [ -d build ]; then rm -rf build; fi


.PHONY: build run clean
