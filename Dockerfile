FROM golang:alpine AS build-stage

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ARG REF

ENV GITHUB_REF_NAME=$REF
ENV DATE=$(date)

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -X github.com/templates/internal/api.defaultVersion=$GITHUB_REF_NAME -X github.com/templates/internal/api.defaultDate=$Date" -o app .

FROM alpine as final

WORKDIR /

COPY --from=build-stage /src/app .

EXPOSE 8080

ENTRYPOINT ["/app"]


