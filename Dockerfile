FROM golang:1.13.0 As goimage
ENV GO111MODULE=on
WORKDIR /go/src/personell/datastructure-algorithem
COPY . /go/src/personell/datastructure-algorithem
#RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig
#RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o datastructure-algorithem main.go

FROM golang:1.13.0
ENV GO111MODULE=on
WORKDIR /datastructure-algorithem/

COPY --from=goimage /go/src/personell/datastructure-algorithem .

CMD ["./datastructure-algorithem"]
