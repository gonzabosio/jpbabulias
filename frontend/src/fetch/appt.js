import { backurl } from "../main"
import { checkCookie } from "../router"
import { deleteCookie, logout } from "./user"

export const getDayAppointments = async (date) => {
    try {
        // console.log(date)
        const resp = await fetch(backurl + '/appointment/day?date=' + date, {
            credentials: 'include'
        })
        const payload = await resp.json()
        // console.log(payload)
        if (!resp.ok) {
            if (resp.status === 401) {
                if (checkCookie('refresh_token')) {

                } else {
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
            }
            console.error(payload.error_dets)
            return { error: true, code: resp.status, message: payload.message }
        }
        return { error: false, code: resp.status, appointments: payload.appointments }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}

export const getFullyBookedDates = async () => {
    try {
        const resp = await fetch(backurl + '/appointment/full', {
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
        return { error: false, code: resp.status, fully_booked_dates: payload.fully_booked_dates }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}