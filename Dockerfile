FROM alpine:3.20
RUN apk add --no-cache bash curl
WORKDIR /app
CMD ["sh","-c","echo 'auth-service-go-postgres container started'; sleep infinity"]
