import { createRouter, createWebHistory } from 'vue-router'
import Landing from '../views/Landing.vue'
import Appointments from '../views/Appointments.vue'
import NotFound from '../views/NotFound.vue'
import Profile from '../views/Profile.vue'
import Treatments from '../views/Treatments.vue'
import ApptConfirmation from '../views/ApptConfirmation.vue'
import Register from '../views/Register.vue'
import { CookiesExist } from '../fetch/user'

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
            component: Appointments,
            meta: {
                requiresAuth: true
            }
        },
        {
            path: '/turnos/confirmar',
            name: 'appointmentConfirmation',
            component: ApptConfirmation,
            meta: {
                requiresAuth: true
            }
        },
        {
            path: '/perfil',
            name: 'profile',
            component: Profile,
            meta: {
                requiresAuth: true
            }
        },
        {
            path: '/tratamientos',
            name: 'treatments',
            component: Treatments
        },
        {
            path: '/registro',
            name: 'register',
            component: Register
        },
        {
            path: '/:pathMatch(.*)*',
            name: 'not-found',
            component: NotFound
        }
    ]
})



let lastAttemptedRoute = null
router.beforeEach(async (to, from, next) => {
    if (to.meta.requiresAuth) {
        const cookiesExist = await CookiesExist()
        lastAttemptedRoute = to.fullPath
        if (!cookiesExist) {
            next({
                path: "/registro",
                query: { mode: 'login' }

            }) // redirect unauthorized users
        } else {
            next() // allow navigation
        }
    } else {
        next()
    }
})

export default router

export function getLastAttemptedRoute() {
    return lastAttemptedRoute
}