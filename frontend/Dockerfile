# Stage 1: Build using updated Node.js
FROM node:20.13.1-alpine3.19 AS build

WORKDIR /app

# 1. Copy package files first
COPY package.json package-lock.json ./

# 2. Install exact versions with clean setup
RUN npm clean-install --no-audit --force

# 3. Copy source files and build
COPY . .
RUN npm run build

# Stage 2: Production image
FROM nginx:1.25.5-alpine3.19
COPY --from=build /app/build /usr/share/nginx/html
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]