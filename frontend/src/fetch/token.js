import { backurl } from "../main"

export const updateAccessToken = async () => {
    const result = await fetch(backurl + '/user/refresh', {
        credentials: 'include'
    })
    const payload = await result.json()
}