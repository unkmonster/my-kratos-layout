FROM golang:1.24.6 AS builder

ARG SERVICE_NAME
ARG GO_MODULE_NAME

# 缓存 go mod download
WORKDIR /src
COPY go.mod go.sum .

RUN go mod download

COPY ./api ./api
COPY ./.git ./.git
COPY ./third_patry ./third_patry
COPY ./app/${SERVICE_NAME} ./app/${SERVICE_NAME}
RUN make -C app/${SERVICE_NAME} build

FROM debian:stable-slim

ARG SERVICE_NAME

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/app/${SERVICE_NAME}/bin/${SERVICE_NAME} /app/server

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

ENV SERVICE_NAME=${SERVICE_NAME}
CMD ["./server", "-conf", "/data/conf"]
