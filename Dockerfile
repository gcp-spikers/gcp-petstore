FROM alpine

RUN ls -al
COPY gopath/bin/gcp-petstore /go/bin/gcp-petstore

ENTRYPOINT /go/bin/gcp-petstore