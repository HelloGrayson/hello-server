FROM golang
ADD . /go/src/github.com/breerly/hello-server
RUN go install github.com/breerly/hello-server
ENTRYPOINT /go/bin/hello-server
EXPOSE 8000-8050
