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
  <TopBar v-if="!isNotFound" />
  <RouterView />
  <Footer v-if="!isNotFound" />
</template>

<style scoped></style>
