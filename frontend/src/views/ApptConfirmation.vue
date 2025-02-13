<script setup>
import { reactive, computed, ref } from 'vue'

const f = new Intl.DateTimeFormat('es-ar', {
    dateStyle: "full",
    timeStyle: "short",
})

const selectedDate = new Date(sessionStorage.getItem('selectedDate'))
const dateToConfirm = f.format(selectedDate)

const formData = reactive({
    fullName: '',
    email: '',
    phone: '',
    dni: '',
    hin: ''
    // agreement: false
})

const isSubmitted = ref(false)

const isFormValid = computed(() => {
    return (
        formData.fullName &&
        formData.email &&
        formData.phone &&
        formData.dni &&
        formData.hin
    )
})

const submitForm = () => {
    if (isFormValid.value) {
        console.log('Form submitted:', formData)
        isSubmitted.value = true
        // send the data to a server
    }
}


</script>

<template>
    <div class="form-container">
        <h2>Confirma tu turno</h2>
        <p>{{ String(dateToConfirm).charAt(0).toUpperCase() + String(dateToConfirm).slice(1) }}</p>
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="fullName">Nombre completo</label>
                <input type="text" id="fullName" v-model="formData.fullName" required />
            </div>
            <div class="form-group">
                <label for="email">Correo</label>
                <input type="email" id="email" v-model="formData.email" required />
            </div>
            <div class="form-group">
                <label for="phone">Número de teléfono</label>
                <input type="tel" id="phone" v-model="formData.phone" required />
            </div>
            <div class="form-group">
                <label for="dni">DNI</label>
                <input type="number" id="dni" v-model="formData.dni" required />
            </div>
            <div class="form-group">
                <label for="hin">Obra social</label>
                <input type="text" id="hin" v-model="formData.hin" required />
            </div>
            <button type="submit" :disabled="!isFormValid">Confirmar</button>
        </form>
        <div v-if="isSubmitted" class="success-message">
            Turno agendado 📆
        </div>
    </div>
</template>

<style scoped>
.form-container {
    max-width: 500px;
    margin: 9em auto;
    padding: 30px;
    background-color: #f5f5f5;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

h2 {
    color: #333;
}

.form-group {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 15px;
}

label {
    display: block;
    margin-bottom: 5px;
    color: #666;
}

/* input[type="date"] */

input[type="text"],
input[type="email"],
input[type="tel"],
input[type="number"] {
    width: 100%;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
    font-size: 16px;
}

button {
    width: 100%;
    padding: 10px;
    background-color: #3790D0;
    font-weight: 600;
    color: white;
    border: none;
    border-radius: 4px;
    font-size: 16px;
    cursor: pointer;

    &:hover {
        background-color: #2176b3;
    }
}

button:disabled {
    background-color: #cccccc;
    cursor: not-allowed;
}

.success-message {
    margin-top: 20px;
    padding: 10px;
    background-color: #dff0d8;
    border: 1px solid #d6e9c6;
    color: #3c763d;
    border-radius: 4px;
    text-align: center;
}
</style>