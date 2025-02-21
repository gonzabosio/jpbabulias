import { backurl } from "../main"
import { checkCookie } from "../router"
import { logout } from "./user"

export const getDayAppointments = async (date, retry = true) => {
    try {
        const resp = await fetch(backurl + '/appointment/day?date=' + date, {
            credentials: 'include'
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 401) {
                if (retry && checkCookie('access_token')) {
                    return await getDayAppointments(date, false)
                } else {
                    const logoutResp = await logout()
                    if (logoutResp.error) {
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

export const getFullyBookedDates = async (retry = true) => {
    try {
        const resp = await fetch(backurl + '/appointment/full', {
            credentials: 'include'
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 401) {
                if (retry && checkCookie('access_token')) {
                    return await getFullyBookedDates(false)
                } else {
                    const logoutResp = await logout()
                    if (logoutResp.error) {
                        return { error: true, code: 401, message: 'Sesión expirada' }
                    }
                    return { error: true, code: 401, message: 'Sesión expirada' }
                }
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

export const saveAppointment = async (apptDate, subject, patientId, email, retry = true) => {
    console.log(apptDate, subject, patientId, email)
    try {
        const resp = await fetch(backurl + '/appointment', {
            method: 'POST',
            credentials: 'include',
            headers: { 'Content-Type': 'application-json' },
            body: JSON.stringify({
                appt_date: apptDate,
                subject: subject,
                patient_id: patientId,
                email: email
            })
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 401) {
                if (retry && checkCookie('access_token')) {
                    return await saveAppointment(apptDate, subject, patientId, retry, false)
                } else {
                    const logoutResp = await logout()
                    if (logoutResp.error) {
                        return { error: true, code: 401, message: 'Sesión expirada' }
                    }
                    return { error: true, code: 401, message: 'Sesión expirada' }
                }
            }
            console.error(payload.error_dets)
            return { error: true, code: resp.status, message: payload.message }
        }
        return { error: false, code: resp.status, appointment: payload.appt_id }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}