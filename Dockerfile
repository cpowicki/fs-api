FROM golang:1.16-alpine
WORKDIR /app
COPY dist/fs-api-linux-amd64 ./fs-api
CMD [ "./fs-api" ]