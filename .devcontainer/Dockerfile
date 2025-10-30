FROM golang:1.25

# System-Basics
RUN apt-get update \
  && apt-get install -y --no-install-recommends \
     build-essential \
     git \
     ca-certificates \
     curl \
     postgresql-client \
     nodejs \
     npm \
     sudo \
  && rm -rf /var/lib/apt/lists/*

# Non-root User anlegen
ARG USERNAME=developer
ARG USER_UID=1000
ARG USER_GID=${USER_UID}
RUN groupadd --gid ${USER_GID} ${USERNAME} \
  && useradd --uid ${USER_UID} --gid ${USER_GID} -m ${USERNAME}

# Arbeitsverzeichnis (optional, hilft bei relativen Pfaden)
WORKDIR /workspaces

# - dlv gotests gopls goimports staticcheck
RUN go install github.com/go-delve/delve/cmd/dlv@latest \
 && go install github.com/cweill/gotests/gotests@v1.6.0 \
 && go install golang.org/x/tools/gopls@latest \
 && go install golang.org/x/tools/cmd/goimports@latest \
 && go install honnef.co/go/tools/cmd/staticcheck@latest




# Non-root User verwenden
USER ${USERNAME}
