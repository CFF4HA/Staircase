FROM nginx:alpine

ARG FRONTEND_URL="http://staircase-webfrontend:8080"
ARG BACKEND_URL="http://staircase-apibackend:8081"

COPY nginx.conf.template /etc/nginx/nginx.conf.template

RUN envsubst '${FRONTEND_URL} ${BACKEND_URL}' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

EXPOSE 80
