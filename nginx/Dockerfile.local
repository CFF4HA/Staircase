FROM nginx:alpine

ARG FRONTEND_URL="http://frontend_container_name:frontend_port"
ARG BACKEND_URL="http://backend_container_name:backend_port"

COPY nginx.conf.template /etc/nginx/nginx.conf.template

RUN envsubst '${FRONTEND_URL} ${BACKEND_URL}' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

EXPOSE 80
