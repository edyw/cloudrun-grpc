FROM node:18-alpine

WORKDIR /app
COPY ./node-api/package*.json ./
RUN npm install
COPY ./node-api/. .

ENTRYPOINT ["node", "index.js"]