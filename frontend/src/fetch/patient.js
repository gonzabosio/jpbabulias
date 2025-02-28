import { backurl } from "../main"
import { logout } from "./user"

export const getPatientsDataByUserId = async (userId) => {
    try {
        const resp = await fetch(backurl + '/patient/' + userId, {
            credentials: 'include'
        })
        const payload = await resp.json()
        // console.log(payload)
        if (!resp.ok) {
            if (resp.status === 401) {
                const tokensExist = await CookiesExist()
                if (tokensExist) {
                    return await getPatientsDataByUserId(userId, false)
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
        return { error: false, code: resp.status, patients: payload.patients }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}

export const savePatient = async (formData, userId) => {
    // console.log(formData, userId)
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
                const tokensExist = await CookiesExist()
                if (tokensExist) {
                    return await savePatient(formData, userId, false)
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
        return { error: false, code: resp.status, patients: payload.patients }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}

export const editPatientData = async (formData, patientId) => {
    try {
        const resp = await fetch(backurl + '/patient', {
            method: 'PATCH',
            credentials: 'include',
            headers: { 'Content-Type': 'application-json' },
            body: JSON.stringify({
                id: patientId,
                first_name: formData.firstName,
                last_name: formData.lastName,
                phone_number: '+54' + formData.phone,
                dni: Number(formData.dni),
                health_insurance: formData.hin,
            })
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 401) {
                const tokensExist = await CookiesExist()
                if (tokensExist) {
                    return await editPatientData(formData, patientId, false)
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
        return { error: false, code: resp.status, message: payload.message }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}

export const deletePatientById = async (patientId) => {
    try {
        const resp = await fetch(backurl + '/patient/' + patientId, {
            method: 'DELETE',
            credentials: 'include'
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 401) {
                const tokensExist = await CookiesExist()
                if (tokensExist) {
                    return await deletePatientById(patientId, false)
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
        return { error: false, code: resp.status, patients: payload.patients }
    } catch (error) {
        console.error(error.message)
        return { error: true, code: 0, message: error.message }
    }
}