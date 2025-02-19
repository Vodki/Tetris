FROM nginx:1.25.5-alpine3.19

# Remove default config
RUN rm /etc/nginx/conf.d/default.conf

# Copy custom NGINX config
COPY nginx/nginx.conf /etc/nginx/conf.d/app.conf

# Certbot SSL certificates (optional)
COPY nginx/certbot/ /etc/letsencrypt/