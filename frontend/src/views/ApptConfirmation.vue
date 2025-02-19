<script setup>
import { reactive, computed, ref, onMounted, onBeforeMount, watch } from 'vue'
import { getPatientsDataByUserId } from '../fetch/patient'
import { useToast } from "vue-toastification";

const toast = useToast()

const localStoredUserDate = JSON.parse(localStorage.getItem("user") || "{}")

const patients = ref([])
const patientSelected = ref({})
const patientSelectedId = ref('')

onBeforeMount(async () => {
    setTimeout(async () => {
        if (!patients.value.length) {
            // show loading spinner component while getting patient data
            const response = await getPatientsDataByUserId(localStoredUserDate.user_id)
            if (response.error) {
                if (result.code === 401) {
                    toast.info(result.message)
                    router.replace({
                        path: '/registro',
                        query: { mode: 'login' }
                    })
                    return
                }
                console.error(response.message)
                toast.error('Error al intentar cargar datos del paciente')
            } else {
                patients.value = response.patients
                const defaultPatient = patients.value.find(opt => opt.main)
                if (defaultPatient) {
                    patientSelected.value = defaultPatient
                    patientSelectedId.value = patientSelected.value.id
                } else {
                    toast.error('Error al intentar cargar datos del paciente')
                }
                console.log('Selected:', patientSelected.value)
            }
        }

    }, 500)
})

watch(patientSelectedId, async (newId) => {
    console.log('search data of:', newId)
    // load data from new patient selected
})

const f = new Intl.DateTimeFormat('es-ar', {
    dateStyle: "full",
    timeStyle: "short",
})

const selectedDate = new Date(sessionStorage.getItem('selectedDate'))
const dateToConfirm = f.format(selectedDate)

const formData = reactive({
    subject: "",
    fullName: "",
    email: localStoredUserDate.email,
    phone: "",
    dni: "",
    hin: ""
})

watch(patientSelected, (loadedPatient) => {
    if (loadedPatient) {
        formData.fullName = loadedPatient.first_name + ' ' + loadedPatient.last_name
        formData.phone = loadedPatient.phone_number
        formData.dni = loadedPatient.dni
        formData.hin = loadedPatient.health_insurance
    }
})

const isSubmitted = ref(false)

const isFormValid = computed(() => {
    return (
        formData.subject &&
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
        <select v-if="patientSelectedId" name="patient" id="patient-selection" v-model="patientSelectedId">
            <option v-for="patient in patients" :key="patient.id" :value="String(patient.id)">
                {{ patient.first_name + ' ' + patient.last_name }}
            </option>
        </select>
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="subject">Asunto (max. 50 caracteres)</label>
                <input type="text" id="subject" v-model="formData.subject" required maxlength="50" />
            </div>
            <div class="form-group">
                <label for="fullName">Nombre completo</label>
                <input type="text" id="fullName" v-model="formData.fullName" disabled />
            </div>
            <div class="form-group">
                <label for="email">Correo</label>
                <input type="email" id="email" :value="formData.email" disabled />
            </div>
            <div class="form-group">
                <label for="phone">Número de teléfono</label>
                <input type="tel" id="phone" v-model="formData.phone" required />
            </div>
            <div class="form-group">
                <label for="dni">DNI</label>
                <input type="number" id="dni" v-model="formData.dni" disabled />
            </div>
            <div class="form-group">
                <label for="hin">Obra social</label>
                <input type="text" id="hin" v-model="formData.hin" disabled />
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
    margin: 5em auto;
    padding: 30px;
    background-color: #f5f5f5;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    padding-top: 6em;
}

h2 {
    color: #333;
}

.form-group {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 1em;

    &:first-child {
        margin-top: 2em;
    }
}

label {
    display: flex;
    width: 85%;
    margin-left: 1em;
    margin-bottom: 0.5em;
    color: #666;
}

#patient-selection {
    width: 50%;
    height: 30px;
    font-size: 16px;
    background-color: white;
    border: 1px solid #ddd;
    border-radius: 0;
}

input[type="text"],
input[type="email"],
input[type="tel"],
input[type="number"] {
    width: 80%;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 0.3em;
    font-size: 16px;
}

button {
    width: 90%;
    padding: 12px;
    margin-top: 2em;
    background-color: #3790D0;
    font-weight: 600;
    color: white;
    border: none;
    border-radius: 0.5em;
    font-size: 1.2rem;
    cursor: pointer;
    transition: 0.25s;

    border: 4px solid transparent;

    &:disabled {
        background-color: gray;

        &:hover {
            border: 4px solid gray;
        }
    }

    &:hover {
        border: 4px solid #2176b3;
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