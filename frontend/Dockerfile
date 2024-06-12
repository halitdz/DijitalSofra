FROM node:20-alpine3.19 as build

WORKDIR /app

COPY ./package.json /app
COPY . .


RUN npm install && npm run dev
