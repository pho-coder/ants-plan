FROM golang
ADD btc-tracker /go/src/github.com/pho-coder/ants-plan/btc-tracker
RUN go get github.com/emicklei/go-restful
RUN go install github.com/pho-coder/ants-plan/btc-tracker
ENTRYPOINT /go/bin/btc-tracker
