FROM nginx:alpine

# Copy the config to the 'templates' directory. 
# Nginx will output the final result to /etc/nginx/nginx.conf automatically.
COPY nginx.conf /etc/nginx/templates/nginx.conf.template

# Default variables if none are provided by Railway
ENV FRONTEND_URL=http://staircase-webfrontend:8080
ENV BACKEND_URL=http://staircase-apibackend:8081

EXPOSE 80
