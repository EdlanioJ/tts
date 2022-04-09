FROM golang:1.18

WORKDIR /go/src

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

COPY . .

RUN apt-get update && \
  apt-get install build-essential protobuf-compiler -y && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 && \
  go install github.com/golang/mock/mockgen@v1.6.0  && \
  wget https://github.com/ktr0731/evans/releases/download/0.10.4/evans_linux_amd64.tar.gz && \
  tar -xzvf evans_linux_amd64.tar.gz && \
  mv evans ../bin && rm -f evans_linux_amd64.tar.gz

CMD ["tail", "-f", "/dev/null"]
