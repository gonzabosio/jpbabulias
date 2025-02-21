import { ref } from 'vue'
import { CookiesExist } from '../../fetch/user'

export const isLoggedIn = ref(false)
export const userData = ref(null)

export const switchAuthStatus = () => {
    const tokensExist = CookiesExist()
    userData.value = JSON.parse(localStorage.getItem('user'))
    if (tokensExist && userData.value) {
        isLoggedIn.value = true
    }
}

export const removeUserData = () => {
    // deleteCookie('refresh_token')
    // deleteCookie('access_token')
    localStorage.removeItem('user')
    userData.value = null
    isLoggedIn.value = false
}
