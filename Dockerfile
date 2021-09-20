FROM golang:1.16

COPY . /go/src

RUN go get golang.org/dl/gotip \
    && gotip download dev.fuzz

WORKDIR /go/src/go-fuzzing-test/types

ENTRYPOINT ["gotip", "test", "-fuzz=FuzzOnlyCertainNumbers" , "-fuzztime=5m"]
