<script setup>
import { RouterView } from 'vue-router';
import TopBar from '../src/components/TopBar.vue'
import Footer from '../src/components/Footer.vue'

import { computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const isNotFound = computed(() => route.name === 'not-found');
</script>

<template>
  <div class="app-container">
    <TopBar v-if="!isNotFound" />
    <main class="main-content">
      <RouterView />
    </main>
    <Footer v-if="!isNotFound" />
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.main-content {
  flex: 1;
  overflow-y: auto;
}
</style>
