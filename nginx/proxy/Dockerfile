FROM nginx:1.23.1-alpine

WORKDIR /home/noednett/Projects/OpenLife/nginx/proxy

COPY ./proxy.conf /etc/nginx/conf.d/openlife-proxy.conf

COPY ./backend-not-found.html /var/www/html/backend-not-found.html

COPY ./includes/ /etc/nginx/includes/

COPY ./ssl/ /etc/ssl/certs/nginx/

COPY ./index.html /var/www/html/test/index.html
