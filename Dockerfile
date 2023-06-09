FROM golang:1.18-alpine as build
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

RUN go build -o /gotypi




## Deploy
FROM alpine:latest as prod

WORKDIR /app
COPY --from=build /gotypi /app/gotypi-app
COPY --from=build /app/templates/ /app/templates
COPY --from=build /app/static/ /app/static
RUN ls -lah /app

EXPOSE 8082

ENTRYPOINT ["/app/gotypi-app"]
#ENTRYPOINT ["tail", "-f", "/dev/null"]


