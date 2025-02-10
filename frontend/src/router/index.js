import { createRouter, createWebHistory } from 'vue-router'
import Landing from '../views/Landing.vue'
import Appointments from '../views/Appointments.vue'
import NotFound from '../views/NotFound.vue'
import Profile from '../views/Profile.vue'
import Treatments from '../views/Treatments.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'landing',
            component: Landing
        },
        {
            path: '/turnos',
            name: 'appointments',
            component: Appointments
        },
        {
            path: '/perfil',
            name: 'profile',
            component: Profile
        },
        {
            path: '/tratamientos',
            name: 'treatments',
            component: Treatments
        },
        {
            path: '/:pathMatch(.*)*',
            name: 'not-found',
            component: NotFound
        }
    ]
})

export default router