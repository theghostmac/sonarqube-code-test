FROM alpine:latest
LABEL maintainer="MacBobby Chibuzor <macbobbychibuzor@gmail.com>"
LABEL description="This is a demo project running an nginx server to explain how to use SonarQube for static code analysis."
RUN apk add --update nginx && \
    rm -rf /var/cache/apk/* \
    mkdir -p /tmp/nginx/
COPY files/nginx.conf /etc/nginx/nginx.conf
COPY files/default.conf /etc/nginx/conf.d/default.conf
ADD files/html.tar.gz /usr/share/nginx/

EXPOSE 80/tcp

ENTRYPOINT ["nginx"]
CMD ["-g", "daemon off;"]