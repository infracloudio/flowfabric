ARG GO_VERSION=1.11.5

FROM golang:${GO_VERSION} AS builder
ENV GOPATH /usr
ENV APP ${GOPATH}/src/github.com/infracloudio/flowfabric/app
COPY /app ${APP}/
WORKDIR ${APP}/cmd/
RUN apt update && apt install -y libpcap-dev
RUN CGO_ENABLED=1 GOOS=linux go build -a -o /app ./server
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /flowctl ./client
#FROM golang:1.11-alpine
#WORKDIR /
#COPY --from=builder /app /
#RUN apk --no-cache add --update libpcap-dev
ENTRYPOINT ["/app"]
