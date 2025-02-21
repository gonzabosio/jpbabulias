import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import VCalendar from 'v-calendar';
import 'v-calendar/style.css';
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

export const backurl = import.meta.env.VITE_BACKEND_URL
export const domain = import.meta.env.DOMAIN

createApp(App)
    .use(router)
    .use(VCalendar)
    .use(Toast)
    .mount('#app')
