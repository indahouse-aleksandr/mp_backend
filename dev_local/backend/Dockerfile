FROM golang:1.21.0-bullseye
ARG USER_ID
ARG GROUP_ID=$USER_ID
ARG USER_NAME
RUN useradd ${USER_NAME} && usermod -u ${USER_ID} ${USER_NAME} && groupmod -g ${GROUP_ID} ${USER_NAME}
USER ${USER_NAME}
WORKDIR /var/www/app
