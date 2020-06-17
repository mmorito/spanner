# ------------------------------
# deployment
# ------------------------------
FROM ubuntu:18.04 as base

# Install libraries
RUN apt-get update && \
  apt-get install -y software-properties-common && \
  apt-add-repository -y "deb http://security.ubuntu.com/ubuntu xenial-security main" && \
  apt-get -yq update && \
  apt-get install -y libxrender1 libfontconfig1 libx11-dev libjpeg62 libxtst6 libssl1.0.0 \
  wget unzip git vim fontconfig pkg-config build-essential xfonts-75dpi xfonts-base libpng12-0 libjpeg-turbo8 \
  libmagickcore-dev libmagickwand-dev


# ------------------------------
# builder
# ------------------------------
FROM base as builder

# installgo lang
RUN apt-get install -y gcc make git
RUN apt-get update
RUN add-apt-repository ppa:longsleep/golang-backports
RUN apt-get install -y golang-1.13

# work directory
RUN mkdir /rest-api
WORKDIR /rest-api

ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:$GOPATH/src/app:/usr/lib/go-1.13/bin:$PATH

ENV GO111MODULE=on

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./bin/app


# ------------------------------
# deployment
# ------------------------------
FROM alpine:3.9 as deployment

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /rest-api/bin/ ./

CMD ["./app"]
EXPOSE 8080
