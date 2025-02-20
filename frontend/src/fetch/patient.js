import { backurl } from "../main"
import { deleteCookie, logout } from "./user"

export const getPatientsDataByUserId = async (userId, retry = true) => {
    try {
        const resp = await fetch(backurl + '/patient/' + userId, {
            credentials: 'include'
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 401) {
                console.log('retry')
                if (retry && checkCookie('refresh_token')) {
                    return await getPatientsDataByUserId(userId, false)
                } else {
                    console.log('trigger logout')
                    const result = await logout()
                    if (result.error) {
                        console.error(result.message)
                        const logoutResp = await logout()
                        if (logoutResp.error) {
                            deleteCookie('access_token')
                            deleteCookie('refresh_token')
                            return { error: true, code: 401, message: 'Sesión expirada' }
                        }
                        return { error: true, code: 401, message: 'Sesión expirada' }
                    }
                    return { error: true, code: 401, message: 'Sesión expirada' }
                }
            }
            console.error(payload.error_dets)
            return { error: true, code: resp.status, message: payload.message }
        }
        return { error: false, code: resp.status, patients: payload.patients }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}

export const savePatient = async (formData, userId, retry = true) => {
    console.log(formData, userId)
    try {
        const resp = await fetch(backurl + '/patient', {
            method: 'POST',
            credentials: 'include',
            headers: { 'Content-Type': 'application-json' },
            body: JSON.stringify({
                first_name: formData.firstName,
                last_name: formData.lastName,
                phone_number: '+54' + formData.phone,
                dni: formData.dni,
                health_insurance: formData.hin,
                user_id: userId
            })
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 401) {
                if (retry && checkCookie('refresh_token')) {
                    return await savePatient(formData, userId, false)
                } else {
                    const result = await logout()
                    if (result.error) {
                        console.error(result.message)
                        const logoutResp = await logout()
                        if (logoutResp.error) {
                            deleteCookie('access_token')
                            deleteCookie('refresh_token')
                            return { error: true, code: 401, message: 'Sesión expirada' }
                        }
                        return { error: true, code: 401, message: 'Sesión expirada' }
                    }
                    return { error: true, code: 401, message: 'Sesión expirada' }
                }
            }
            console.error(payload.error_dets)
            return { error: true, code: resp.status, message: payload.message }
        }
        return { error: false, code: resp.status, patients: payload.patients }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}