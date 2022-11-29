FROM alpine:3.17.0

RUN apk update && apk add --no-cache ca-certificates && \
    apk add tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

ADD ./bin/blog2 /go/bin/blog2
ADD ./config /go/src/blog2/config
ADD ./build /go/src/blog2/build
ADD ./pages /go/src/blog2/pages

WORKDIR /go/src/blog2
CMD ["/go/bin/blog2"]
EXPOSE 8080
