<script setup>
import { RouterView } from 'vue-router';
import TopBar from '../src/components/TopBar.vue'
import Footer from '../src/components/Footer.vue'
import Chatbot from '../src/Components/Chatbot.vue'

import { computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const dontShow = computed(() => route.name === 'not-found' || route.name === 'register');
</script>

<template>
  <div class="app-container">
    <TopBar v-if="!dontShow" />
    <main class="main-content">
      <RouterView />
      <Chatbot v-if="!dontShow" />
    </main>
    <Footer v-if="!dontShow" />
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
