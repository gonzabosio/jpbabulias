import { backurl } from "../main"
import { deleteCookie, logout } from "./user"

export const getPatientsDataByUserId = async (userId) => {
    try {
        const resp = await fetch(backurl + '/patient/' + userId, {
            credentials: 'include'
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 401) {
                const result = await logout()
                if (result.error) {
                    console.error(result.message)
                    const secTry = await logout()
                    if (secTry.error) {
                        deleteCookie('access_token')
                        deleteCookie('refresh_token')
                        return { error: true, code: 401, message: 'Sesión expirada' }
                    }
                    return { error: true, code: 401, message: 'Sesión expirada' }
                }
                return { error: true, code: 401, message: 'Sesión expirada' }
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