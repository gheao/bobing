FROM golang
RUN set -ex \
&& go env -w GO111MODULE=on \
&& go env -w GOPROXY=https://goproxy.cn,direct
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o go go.go
EXPOSE 8008
ENTRYPOINT ["./go"]