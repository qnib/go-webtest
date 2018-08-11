FROM golang AS build

WORKDIR /go/src/github.com/qnib/go-webtest/
COPY main.go .
COPY ./lib ./lib
COPY ./vendor ./vendor
RUN go build -o webtest -ldflags "-linkmode external -extldflags -static" -a main.go
RUN ./webtest --version

FROM scratch
ENV WEBTEST_HTTP_HOST="0.0.0.0"
ENV WEBTEST_HTTP_PORT="8081"
COPY --from=build /go/src/github.com/qnib/go-webtest/webtest /webtest
CMD ["/webtest"]
