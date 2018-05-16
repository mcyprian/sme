FROM golang:1.10.2

LABEL maintainer="Michal Cyprian <mcyprian@mail.muni.cz>"

ENV SMEPATH /go/src/github.com/mcyprian/sme

EXPOSE 8080

COPY . $SMEPATH

RUN cd $SMEPATH && go get && make

WORKDIR $SMEPATH

CMD ["./sme"]
