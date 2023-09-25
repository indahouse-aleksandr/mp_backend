docker build --tag 'prod_mp_backend' --platform=linux/amd64 . && docker run --name 'prod_mp_backend' --rm -p 80:80 'prod_mp_backend'
