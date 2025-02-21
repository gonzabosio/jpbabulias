import { ref } from 'vue'
import { checkCookie } from '../../router'
import { deleteCookie } from '../../fetch/user'

export const isLoggedIn = ref(false)
export const userData = ref(null)

export const switchAuthStatus = () => {
    const tokenExist = checkCookie('refresh_token')
    userData.value = JSON.parse(localStorage.getItem('user'))
    if (tokenExist && userData.value) {
        isLoggedIn.value = true
    }
}

export const removeUserData = () => {
    deleteCookie('refresh_token')
    deleteCookie('access_token')
    localStorage.removeItem('user')
    userData.value = null
    isLoggedIn.value = false
}
