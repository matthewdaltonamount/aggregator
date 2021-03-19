FROM golang:1.14.1-stretch
RUN apt-get install ca-certificates
COPY ./certs /usr/local/share/ca-certificates/
RUN update-ca-certificates
ADD . /
WORKDIR /
RUN go build -o main .
EXPOSE 3000
CMD ["/main"]
