FROM alpine:latest

WORKDIR "/opt"

ADD .docker_build/heroku-goservice /opt/bin/heroku-goservice

CMD ["/opt/bin/heroku-goservice"]