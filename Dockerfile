FROM golang:1.19.4-alpine3.17 as build
WORKDIR /src
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o keda-app-demo

FROM alpine:3.17.0
COPY --from=build /src/keda-app-demo /
EXPOSE 8080
CMD ["/keda-app-demo"]