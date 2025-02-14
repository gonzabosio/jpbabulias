const backurl = import.meta.env.VITE_BACKEND_URL

export const getDayAppointments = async (date) => {
    try {
        // console.log(date)
        const resp = await fetch(backurl + '/appointment/day?date=' + date)
        const payload = await resp.json()
        // console.log(payload)
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, message: payload.message }
        }
        return { error: false, appointments: payload.appointments }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
    }
}

export const getFullyBookedDates = async () => {
    try {
        const resp = await fetch(backurl + '/appointment/full')
        const payload = await resp.json()
        if (!resp.ok) {
            console.error(payload.error_dets)
            return { error: true, message: payload.message }
        }
        return { error: false, fully_booked_dates: payload.fully_booked_dates }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
    }
}