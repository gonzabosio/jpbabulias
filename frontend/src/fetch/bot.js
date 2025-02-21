import { backurl } from "../main"

export const sendPrompt = async (promptMsg) => {
    try {
        const resp = await fetch(backurl + '/bot/prompt', {
            method: 'POST',
            body: JSON.stringify({
                "prompt": promptMsg
            }),
            headers: {
                'Content-Type': 'application-json'
            }
        })
        const payload = await resp.json()
        if (!resp.ok) {
            if (resp.status === 422) {
                return { error: false, response: payload.response }
            } else {
                console.error(payload.error_dets)
                return { error: true, message: payload.message }
            }
        }
        return { error: false, response: payload.response }
    } catch (error) {
        console.error(error.message)
        return { error: true, message: error.message }
    }
}