<template>
  <main class="main">
    <div class="app-container">
      <div class="tiSb">
      <h1>WallpaperFinder</h1>
        <a class="git" href="https://github.com/AntonioR10"><i class="fa-brands fa-github"></i></a>
      <form @submit.prevent="resetAndFetch" class="search-bar">
        <input 
          type="search" 
          name="search" 
          pattern=".*\S.*" 
          v-model="searchQuery" 
          @keyup.enter="resetAndFetch" 
          required 
        />
        <button type="submit" class="search-btn"><span>Search</span></button>
      </form>
      </div>
      <div v-if="loading" class="loader">Caricamento in corso...</div>

      <div v-else class="gallery">
        <div v-for="photo in wallpapers" :key="photo.id" class="card">
          <img :src="photo.src.medium" :alt="photo.photographer" loading="lazy" />
          <div class="figcaption">
            <p>Autore: {{ photo.photographer }}</p>
            <a :href="photo.url" target="_blank" rel="noopener">Download</a>
          </div>
        </div>
      </div>
    </div>
  </main>
  
  <div class="pages">
    <a @click="changePage(currentPage - 1)" :class="{ disabled: currentPage <= 1 }">
      <i class="fa-solid fa-arrow-left"></i>
    </a>
    <span class="page-number">{{ currentPage }}</span>
    <a @click="changePage(currentPage + 1)">
      <i class="fa-solid fa-arrow-right"></i>
    </a>
  </div>
  <footer class="footer">
    <p>Powered by Pexels - Antonio R.</p>
  </footer>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const searchQuery = ref('')
const wallpapers = ref([])
const loading = ref(false)
const currentPage = ref(1)

const fetchWallpapers = async () => {
  loading.value = true
  try {
    // Mandiamo sia la query che la pagina al nostro backend in Go su Vercel
    const response = await fetch(`/api/index?query=${searchQuery.value}&page=${currentPage.value}`)
    const data = await response.json()
    
    wallpapers.value = data.photos || []
  } catch (error) {
    console.error("Errore nel caricamento dei wallpaper:", error)
  } finally {
    loading.value = false
  }
}

// Cambia pagina e ricarica
const changePage = (newPage) => {
  if (newPage < 1) return
  currentPage.value = newPage
  fetchWallpapers()
}

// Se l'utente fa una nuova ricerca, resetta alla pagina 1
const resetAndFetch = () => {
  currentPage.value = 1
  fetchWallpapers()
}

onMounted(() => {
  fetchWallpapers()
})
</script>

<style scoped src="./app.css"></style>