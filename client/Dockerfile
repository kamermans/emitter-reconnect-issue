FROM golang:1.12.3

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && go get -u golang.org/x/tools/cmd/goimports \
    && go get -u github.com/emitter-io/emitter

COPY ./ /go/src/emitter-test

RUN cd /go/src/emitter-test \
    && dep ensure -v \
    && go install -v ./... \
    && ls -al $GOPATH/bin
