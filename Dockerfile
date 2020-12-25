FROM golang:latest as goBuilder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .



FROM node:latest as nodeBuilder

WORKDIR /app

COPY ui/web/package.json ui/web/package-lock.json ./

RUN npm install

COPY ./ui/web .

# This will have the api point to itself
ENV NEXT_PUBLIC_API_URL=

RUN npm run export_simple



FROM alpine:latest 

ENV SERVE_DIR=/root/out

ENV SERVE_STATIC=TRUE

ENV ALLOW_ORIGIN_URL=http://10.0.0.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=goBuilder /app/Limitless-Lottery .

COPY --from=nodeBuilder /app/out ./out

EXPOSE 8080

CMD ["./Limitless-Lottery"]