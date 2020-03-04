FROM golang

ADD . /go/src/github.com/deBroglieeeen/fellows

WORKDIR /go/src/github.com/deBroglieeeen/fellows

RUN go get github.com/lib/pq
RUN go install github.com/deBroglieeeen/fellows
# oxequa/realize のライブラリを取ってくる
#RUN go get github.com/oxequa/realize
#RUN go get github.com/pilu/fresh

#CMD ["fresh"]

ENTRYPOINT /go/bin/fellows

EXPOSE 8080