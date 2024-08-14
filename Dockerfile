FROM node:18 AS builder

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY ./backend /app/backend
COPY ./frontend /app/frontend

WORKDIR /app/backend

RUN npm run build

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/backend .

CMD ["node", "main.js"]

EXPOSE 8080