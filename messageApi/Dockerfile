FROM golang:1.22.2

WORKDIR /app
ADD . /app

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main