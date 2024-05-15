FROM golang:latest
 
 
ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/MaidSG/fast-random-go
COPY . $GOPATH/src/github.com/MaidSG/fast-random-go
RUN go build .
 
 
EXPOSE 8089
ENTRYPOINT ["./example.com.web-service-gin"]

#
# Build
# docker build -t go-gin-example .
#
# Run
# docker run -p 8000:8000 go-gin-example
