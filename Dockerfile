FROM golang:1.18-alpine
WORKDIR /app

# we are copying the whole dir
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

# replace .env
RUN rm .env
RUN mv .env.docker .env

# make sure we are in correct directory
RUN ls -lah

#that is for build
#RUN go build -o /docker-gs-ping

EXPOSE 8080
CMD ["./entrypoint.sh"]


