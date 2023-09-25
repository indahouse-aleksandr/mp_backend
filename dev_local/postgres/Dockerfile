FROM postgres:15.4-bullseye
RUN apt update && apt install pgbouncer -y
USER postgres
WORKDIR /var/www/app