const backurl = import.meta.env.VITE_BACKEND_URL

export const getPatientsDataByUserId = async (userId) => {
    try {
        const resp = await fetch(backurl + '/patient/' + userId)
        const payload = await resp.json()
        // console.log('Response:', payload)
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, message: payload.message }
        }
        return { error: false, patients: payload.patients }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
    }
}