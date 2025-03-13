import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  server: {
    host: '0.0.0.0', // Allow external access
    port: 3000,
    strictPort: true,
    hmr: {
      clientPort: 3000, // Important for Docker HMR
    },
  },
  plugins: [
    tailwindcss(),
  ],
})