FROM golang:latest
ENV GO111MODULE=on
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY ./main.go .
RUN go mod init github.com/timtarusov/mux_test && go get github.com/gorilla/mux && go get -u github.com/lib/pq
RUN go build ../app
EXPOSE 80
CMD ["go", "run", "../app"]
