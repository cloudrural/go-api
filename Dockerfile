FROM golang:1.16

ENV APP_HOME=/opt/app

COPY . ${APP_HOME}/

WORKDIR ${APP_HOME}

RUN go build

EXPOSE 8080

CMD [ "./go-api" ]
