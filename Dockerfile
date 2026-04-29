FROM alpinelinux/golang:latest AS builder

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod download

COPY ./src/ ./src/
RUN CGO_ENABLED=0 GOOS=linux go build -o ./dist/main ./src/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /bin/

COPY --from=builder /app/dist/main ./

ARG MONGO_URI
ENV GIN_MODE=release
ENV MONGO_URI=${MONGO_URI}
ENV PORT=3000

EXPOSE 3000

CMD [ "./main" ]
