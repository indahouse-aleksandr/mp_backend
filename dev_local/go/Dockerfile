FROM golang:1.21.1-bookworm
ARG USER_ID
ARG GROUP_ID=$USER_ID
RUN useradd -m duser && usermod -u ${USER_ID} duser && groupmod -g ${GROUP_ID} duser
USER duser
WORKDIR /var/www/app
