import { backurl } from '../main'

export const userSignUp = async (formData) => {
    try {
        const resp = await fetch(backurl + '/user/signup',
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application-json'
                },
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
            return { error: true, message: payload.message }
        }
        return { error: false, user_data: payload.user_data }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
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
                body: JSON.stringify({
                    'email': email,
                    'password': password,
                }),
            }
        )
        const payload = await resp.json()
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, message: payload.message }
        }
        return { error: false, user_data: payload.user_data }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
    }
}