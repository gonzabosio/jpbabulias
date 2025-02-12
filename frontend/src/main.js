import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import VCalendar from 'v-calendar';
import 'v-calendar/style.css';
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

createApp(App)
    .use(router)
    .use(VCalendar)
    .use(Toast)
    .mount('#app')
