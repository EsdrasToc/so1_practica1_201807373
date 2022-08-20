#FROM golang:1.16-alpine

#WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./

#RUN go mod download

#COPY . .

#RUN go build -o /back

#EXPOSE 3030

#CMD ["/back"]

FROM golang:1.16-alpine
WORKDIR /go/src/app

COPY . .

RUN go mod download
EXPOSE 3030

CMD ["go", "run", "main.go"]