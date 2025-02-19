<script setup>
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { userLogin, userSignUp } from '../fetch/user'
import { useToast } from "vue-toastification";
import { getLastAttemptedRoute } from '../router';

const toast = useToast()
const route = useRoute()
const router = useRouter()
const formMode = ref(route.query.mode)
const formData = {
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    birthDate: '',
    phoneNumber: '',
    dni: '',
    insurance: ''
}

watch(formMode, () => {
    router.replace({ query: { mode: formMode.value } })
})

const handleLogin = async () => {
    const resp = await userLogin(formData.email, formData.password)
    if (resp.error) {
        toast.error(resp.message)
    } else {
        const userData = {
            user_id: resp.user_data.id,
            email: resp.user_data.email,
            admin: resp.user_data.admin
        }
        localStorage.setItem('user', JSON.stringify(userData))
        const redirect = getLastAttemptedRoute() || '/'
        router.replace(redirect)
    }
}

const handleSignUp = async () => {
    const date = new Date(formData.birthDate)
    formData.birthDate = date.toISOString()
    const resp = await userSignUp(formData)
    if (resp.error) {
        toast.error('Falló el registro. Revise los campos e inténtelo de nuevo')
    } else {
        const userData = {
            user_id: resp.user_data.id,
            email: resp.user_data.email,
            admin: resp.user_data.admin
        }
        localStorage.setItem('user', JSON.stringify(userData))
        const redirect = getLastAttemptedRoute() || '/'
        router.replace(redirect)
    }
}

</script>

<template>
    <div v-if="formMode === 'login'" class="container">
        <div class="form-container">
            <h1>Iniciar sesión</h1>
            <form @submit.prevent="handleLogin">
                <div class="form-group">
                    <label for="email">Correo electrónico</label>
                    <input type="email" id="email" v-model="formData.email" required>
                </div>
                <div class="form-group">
                    <label for="password">Contraseña</label>
                    <input type="password" id="password" v-model="formData.password" required>
                </div>
                <a href="/recuperar" id="recover-pw-link">Olvidé mi contraseña</a>
                <button type="submit" class="submit-btn">
                    Iniciar sesión
                </button>
            </form>
            <button @click="formMode = 'signup'" class="btn-switch-form">Crear cuenta</button>
        </div>
    </div>
    <div v-else class="container">
        <div class="form-container">
            <h1>Registrarse</h1>
            <form @submit.prevent="handleSignUp">
                <div class="form-group">
                    <label for="name">Nombre</label>
                    <input type="text" id="name" v-model="formData.firstName" required>
                </div>
                <div class="form-group">
                    <label for="surname">Apellido</label>
                    <input type="text" id="surname" v-model="formData.lastName" required>
                </div>
                <div class="form-group">
                    <label for="email">Correo electrónico</label>
                    <input type="email" id="email" v-model="formData.email" required>
                </div>
                <div class="form-group">
                    <label for="password">Contraseña</label>
                    <input type="password" id="password" v-model="formData.password" required>
                </div>
                <div class="form-group">
                    <label for="birthdate">Fecha de nacimiento</label>
                    <input type="date" id="birthdate" v-model="formData.birthDate" required>
                </div>
                <div class="form-group">
                    <label for="phone-number">Número de teléfono</label>
                    <div class="phone-input-container">
                        <span class="phone-prefix">+54</span>
                        <input type="tel" id="phone-number" v-model="formData.phoneNumber" placeholder="1234-101010"
                            required>
                    </div>
                </div>
                <div class="form-group">
                    <label for="dni">DNI</label>
                    <input type="number" id="dni" v-model="formData.dni" required>
                </div>
                <div class="form-group">
                    <label for="insurance">Obra Social</label>
                    <input type="text" id="insurance" v-model="formData.insurance" required>
                </div>
                <button type="submit" class="submit-btn">
                    Registrarse
                </button>
            </form>
            <button @click="formMode = 'login'" class="btn-switch-form">Tengo una cuenta</button>
        </div>
    </div>
</template>

<style scoped>
* {
    box-sizing: border-box;
}

h1 {
    margin: 0;
    margin-bottom: 8px;
}

.container {
    display: flex;
    justify-content: center;
    align-items: center;
    overflow: hidden;
    min-height: 100vh;
    background-image: url('../assets/edge-bg.jpg');
    background-size: cover;
}

.form-container {
    background-color: #fff;
    width: 100%;
    max-width: 400px;
    padding: 20px;
    border-radius: 0.5em;
}

.form-group {
    margin-bottom: 1rem;
}

label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
}

input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 0.5em;
    font-size: 16px;
}

.phone-input-container {
    display: flex;
    align-items: center;
    border: 1px solid #ddd;
    border-radius: 0.5em;
    overflow: hidden;
}

.phone-prefix {
    background-color: #f0f2f5;
    color: #555;
    border-right: 1px solid #ddd;
}

.phone-input-container input {
    border: none;
    flex-grow: 1;
}

.submit-btn {
    width: 100%;
    padding: 0.75rem;
    margin-top: 1em;
    background-color: #3790D0;
    color: #fff;
    border: none;
    border-radius: 0.5em;
    font-size: 1rem;
    cursor: pointer;
    font-weight: 600;
    transition: background-color 0.3s ease;
}

.submit-btn:hover {
    background-color: #2176b3;
}

#recover-pw-link {
    cursor: pointer;
    text-decoration: none;
    border: none;
    color: gray;
    font-size: 1rem;
    font-weight: 500;

    &:hover {
        text-decoration: underline;
    }
}

.phone-input-container {
    display: flex;
    align-items: center;
    border: 1px solid #ddd;
    border-radius: 0.5em;
    overflow: hidden;
}

.phone-prefix {
    background-color: #f0f2f5;
    padding: 8px;
    font-size: 16px;
    color: #555;
    border-right: 1px solid #ddd;
}

.phone-input-container input {
    border: none;
    flex-grow: 1;
}

.btn-switch-form {
    cursor: pointer;
    background-color: transparent;
    border: none;
    color: #3790D0;
    margin: 1em 0;
    font-size: 1rem;
    font-family: poppins-regular;

    &:hover {
        text-decoration: underline;
    }
}

@media (max-width: 480px) {
    .form-container {
        padding: 1.5rem;
    }

    h2 {
        font-size: 1.5rem;
    }

    input,
    .submit-btn {
        font-size: 0.9rem;
    }
}
</style>