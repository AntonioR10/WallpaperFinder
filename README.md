# WallpaperFinder

A responsive, modern serverless web application designed to explore and download wallpapers, powered by the official **Pexels API**.

This project leverages **Vue 3** with `<script setup>` syntax for a fluid user interface, combined with a lightweight serverless backend written in **Go (Golang)**, optimized for immediate deployment on **Vercel**.

---

## Tech Stack

* **Frontend:** Vue 3 (Vite, Composition API)
* **Backend:** Go 1.x (Vercel Serverless Functions)
* **Core API:** Pexels API
* **Icons:** Font Awesome 6 (Regular & Solid via CDN)

---

## Local Installation

To run this project locally on your machine for testing or development, follow these steps:

1. **Clone the repository:**
   ```bash
   git clone [https://github.com/AntonioR10/WallpaperFinder.git](https://github.com/AntonioR10/WallpaperFinder.git)
   cd WallpaperFinder
2. **Install frontend dependencies:**
    npm install
3. **Configure the Pexels API Key:**

    - For local testing, add your temporary token to the localPexelsKey variable inside src/App.vue.

    - For production, set the environment variable PEXELS_API_KEY directly in your Vercel dashboard.
4. **Start the development server:**
    npm run dev

## Deploying to Vercel
This project is fully pre-configured to be deployed on Vercel out of the box using the vercel.json architecture and the api/ directory:

 1. Connect your GitHub repository to Vercel.
 2. In your Vercel Project Settings, add the following Environment Variable:
    - Key: PEXELS_API_KEY
    - Value: Your secret Pexels token
 3. Trigger the deployment. The Go function inside /api/index.go will automatically  build as a serverless endpoint.