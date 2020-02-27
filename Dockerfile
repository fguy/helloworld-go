FROM idock.daumkakao.io/driver/go-grpc-base

ARG service=helloworld-go
ARG org=travis-kang
ARG pkg=github.kakaocorp.com/${org}/${service}

ADD . $GOPATH/src/${pkg}

WORKDIR $GOPATH/src/${pkg}

RUN dep ensure

# Install protoc-gen-go, protoc-gen-grpc-gateway, protoc-gen-swagger
RUN go install ./vendor/github.com/golang/protobuf/protoc-gen-go/ \
    ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

EXPOSE 50051

RUN make test && make
CMD ["./server"]
