FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod .
COPY . .
RUN go mod tidy
#RUN go build

ENV ELASTIC_APM_ACTIVE=true
ENV ELASTIC_APM_SERVICE_NAME=apm-test-server
ENV ELASTIC_APM_SECRET_TOKEN=apm-secret-here
ENV ELASTIC_APM_SERVER_URL=apm-server-here
ENV ELASTIC_APM_VERIFY_SERVER_CERT=false
ENV ELASTIC_APM_ENVIRONMENT=local
ENV ELASTIC_APM_CAPTURE_BODY=all
ENV ELASTIC_APM_CLOUD_PROVIDER=none

CMD ["go","run","test-gin-app.go"]
