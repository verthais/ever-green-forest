FROM golang:1.20.1 as dev-swisscom-game

RUN apt update -y && apt upgrade -y

#
# Add dev user with no root rights
ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=$USER_UID

#
# Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd -s /bin/bash --uid $USER_UID --gid $USER_GID -m $USERNAME

# Set home dir
ENV HOME /home/$USERNAME

# Clean up
RUN apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# the user we're applying this too (otherwise it most likely install for root)
USER $USERNAME

RUN go install -v golang.org/x/tools/gopls@latest && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest && \
    go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949

ENTRYPOINT /bin/bash
