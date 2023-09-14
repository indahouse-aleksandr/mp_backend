FROM golang:1.21.0-alpine3.17 AS BuildStage
WORKDIR /var/www/app
COPY . .
RUN go build -o /var/www/app/main ./main.go

FROM scratch
WORKDIR /var/www/app
COPY --from=BuildStage /var/www/app/ /var/www/app/
EXPOSE 80 80
ENV PORT=80
ENTRYPOINT ["/var/www/app/main"]