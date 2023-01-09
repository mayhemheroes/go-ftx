FROM golang:1.18 as builder

COPY . /go-ftx
WORKDIR /go-ftx/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang


RUN go-fuzz-build -libfuzzer -o fuzzAuthAndTypes.a && \
    clang -fsanitize=fuzzer fuzzAuthAndTypes.a -o fuzzAuthAndTypes.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /go-ftx/mayhem/fuzzAuthAndTypes.libfuzzer /