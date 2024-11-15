FROM golang:1.23.3-bookworm

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        sudo \
        curl \
        git \
        fzf \
        bash \
        zsh \
        neovim \
        make \
        golang \
        libgit2-dev

RUN useradd -ms /bin/bash promptorium && \
    echo "promptorium ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

RUN echo "export USER=promptorium" >> /home/promptorium/.bashrc
RUN echo "export USER=promptorium" >> /home/promptorium/.zshrc

USER promptorium
RUN curl https://raw.githubusercontent.com/Promptorium/promptorium/refs/heads/main/install.bash | bash

WORKDIR /home/promptorium
