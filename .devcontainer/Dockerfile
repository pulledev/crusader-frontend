FROM golang:1.25

RUN apt-get update && apt-get install -y build-essential


# PostgreSQL-Client installieren
RUN apt-get install -y postgresql-client

#nodejs and npm
RUN apt-get install nodejs npm -y


ARG USERNAME=developer
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN groupadd --gid $USER_GID $USERNAME && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME
USER $USERNAME