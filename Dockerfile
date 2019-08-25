FROM iron/go:dev

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/gokadin/ai-simplest-network

ADD ./src $SRC_DIR

RUN cd $SRC_DIR; go build -o app; cp app /app/

ENTRYPOINT ["./app"]