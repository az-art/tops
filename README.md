#Tops sidecar container [![CircleCI](https://circleci.com/gh/az-art/tops.svg?style=shield)](https://circleci.com/gh/az-art/tops)

 * Start nginxdemo container e.g. `docker run -p 8080:80 -d nginxdemos/hello`
 * Start sidecar 'tops' `docker run --rm --pid=container:[APP_ID] -p 8000:8000 azzart/tops`

--