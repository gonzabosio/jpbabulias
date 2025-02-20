import { backurl, domain } from '../main'

export const userSignUp = async (formData) => {
    try {
        const resp = await fetch(backurl + '/user/signup',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application-json',
                },
                credentials: 'include',
                body: JSON.stringify({
                    'first_name': formData.firstName,
                    'last_name': formData.lastName,
                    'email': formData.email,
                    'password': formData.password,
                    'birth_date': formData.birthDate,
                    'phone_number': '+54' + formData.phoneNumber,
                    'dni': formData.dni,
                    'health_insurance': formData.insurance
                }),
            }
        )
        const payload = await resp.json()
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, code: resp.status, message: payload.message }
        }
        return { error: false, code: resp.status, user_data: payload.user_data }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}

export const userLogin = async (email, password) => {
    try {
        const resp = await fetch(backurl + '/user/login',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application-json'
                },
                credentials: 'include',
                body: JSON.stringify({
                    'email': email,
                    'password': password,
                }),
            }
        )
        const payload = await resp.json()
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, code: resp.status, message: payload.message }
        }
        return { error: false, code: resp.status, user_data: payload.user_data }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}

export const logout = async (retry = true) => {
    try {
        const resp = await fetch(backurl + '/user/logout', {
            method: 'POST',
            credentials: 'include'
        })
        if (!resp.ok) {
            console.error('Failed to logout')
            if (retry) {
                return await logout(false)
            } else {
                deleteCookie('refresh_token')
                deleteCookie('access_token')
                return { error: false, code: resp.status, message: 'Sesión cerrada' }
            }
        }
        return { error: false, code: resp.status, message: 'Sesión cerrada' }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}

export const deleteCookie = (name) => {
    document.cookie = name + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/; domain=" + domain
}