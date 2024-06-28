FROM golang:1.21 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

FROM build-stage as test
RUN go test -v ./...

FROM scratch

COPY --from=build /entrypoint /entrypoint

EXPOSE 8080

ENTRYPOINT ["/entrypoint"]
