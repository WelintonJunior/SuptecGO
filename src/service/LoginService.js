export class LoginService {
    async Login(e) {
        e.preventDefault()
        const fd = new FormData(e.target)
        const data = Object.fromEntries(fd.entries())
        const response = await fetch("/Login", ({
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                USU_email: data.email,
                USU_password: data.password
            })
        }))
        const result = await response.json()
        return result
    }
}