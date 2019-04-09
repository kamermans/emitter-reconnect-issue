FROM golang:1.12.3

RUN apt-get update \
    && apt-get install -y supervisor \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && go get -u golang.org/x/tools/cmd/goimports \
    && go get -u github.com/emitter-io/emitter

COPY ./supervisor.conf /etc/supervisor/conf.d/supervisor.conf

CMD supervisord
