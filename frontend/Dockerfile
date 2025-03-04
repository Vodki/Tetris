# ========== DEVELOPMENT ========= 

FROM node:22-alpine AS development

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

# Development server with HMR
CMD ["npm", "run", "dev"]

# ========== PRODUCTION =========

FROM node:22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy package files first to cache dependencies
COPY package*.json ./
COPY tsconfig*.json ./
COPY vite.config.ts ./

# Install dependencies
RUN npm install

# Copy all other source files
COPY . .

# Build the application
ENV NODE_ENV=production
RUN npm run build

# Stage 2: Production server
FROM nginx:alpine

# Copy nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy built assets from builder
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port
EXPOSE 80

# Start nginx server
CMD ["nginx", "-g", "daemon off;"]