FROM golang:1.20 as build

WORKDIR /app
COPY . ./
RUN go mod download
WORKDIR /app/go-contact

RUN CGO_ENABLED=0 GOARCH=amd64 go build -mod=readonly -v -o /app/server

FROM gcr.io/distroless/static:nonroot as runtime

COPY --chown=nonroot:nonroot --from=build /app/server /server

ENTRYPOINT ["/server", "-alsologtostderr"]