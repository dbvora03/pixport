FROM golang:alpine
RUN mkdir /app  
ADD ./app /app
WORKDIR /app
RUN go build -o svc .
CMD ["/app/svc"]
