FROM node:latest as builder

WORKDIR /frontend

COPY . .
RUN yarn
RUN yarn run build
RUN ls
RUN pwd

FROM nginx:stable-alpine

COPY --from=builder /frontend/build/ /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]
