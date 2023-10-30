FROM node:16-alpine

COPY package.json package-lock.json ./

RUN npm ci --production

COPY index.js ./

CMD ["node", "index.js"]
