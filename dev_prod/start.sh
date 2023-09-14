docker build --tag 'prod_tc_backend' --platform=linux/amd64 . && docker run --name 'prod_tc_backend' --rm -p 80:80 'prod_tc_backend'
