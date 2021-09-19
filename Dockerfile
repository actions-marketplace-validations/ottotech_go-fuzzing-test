FROM golang:1.16

COPY . /go/src

RUN gotip download dev.fuzz \
    && go get golang.org/dl/gotip

WORKDIR /go/src/go-fuzzing-test/types

CMD ["gotip", "test", "-fuzz=FuzzOnlyCertainNumbers" , "-fuzztime=60s"]
