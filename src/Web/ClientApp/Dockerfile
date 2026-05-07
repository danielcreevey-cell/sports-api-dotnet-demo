FROM node:24 AS build

WORKDIR /app

COPY package.json package.json
COPY package-lock.json package-lock.json

RUN npm ci --ignore-scripts

COPY src/ src/
COPY public/ public/
COPY angular.json tsconfig*.json default.conf.template ./

RUN npm run build

FROM nginx:alpine

COPY --from=build /app/default.conf.template /etc/nginx/templates/default.conf.template
COPY --from=build /app/dist/weather/browser /usr/share/nginx/html

RUN addgroup -g 1001 -S appgroup && adduser -u 1001 -S appuser -G appgroup
USER appuser

# Expose the default nginx port
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]