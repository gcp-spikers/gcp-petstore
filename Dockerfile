FROM alpine

# Copy our binary from the default cloudbuild folder 
COPY gopath/bin/gcp-petstore /go/bin/gcp-petstore

ENTRYPOINT /go/bin/gcp-petstore