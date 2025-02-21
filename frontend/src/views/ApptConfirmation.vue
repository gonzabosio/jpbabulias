<script setup>
import { reactive, computed, ref, onMounted, onBeforeMount, watch } from 'vue'
import { getPatientsDataByUserId } from '../fetch/patient'
import { useToast } from "vue-toastification";
import { saveAppointment } from '../fetch/appt';
import { useRouter } from 'vue-router';

const toast = useToast()
const router = useRouter()
const localStoredUserData = JSON.parse(localStorage.getItem("user") || {})

const patientsList = ref([])
const patientSelected = ref({})
const indexSelected = ref(0)
onBeforeMount(async () => {
    window.scrollTo(0, 0)
    setTimeout(async () => {
        if (!patientsList.value.length) {
            const response = await getPatientsDataByUserId(localStoredUserData.user_id)
            console.log(response)
            if (response.error) {
                if (response.code === 401) {
                    toast.info(response.message)
                    router.replace({
                        path: '/registro',
                        query: { mode: 'login' }
                    })
                    return
                }
                console.error(response.message)
                toast.error('Error al intentar cargar datos del paciente')
            } else {
                patientsList.value = response.patients
                patientSelected.value = patientsList.value.find(opt => opt.patient.main)
                if (!patientsList.value) {
                    toast.error('Error al intentar cargar datos del paciente')
                }
            }
        }
    }, 500)
})

watch(indexSelected, async (newIdx) => {
    patientSelected.value = patientsList.value[newIdx]
})

const f = new Intl.DateTimeFormat('es-ar', {
    dateStyle: "full",
    timeStyle: "short",
})

const selectedDate = new Date(sessionStorage.getItem('selectedDate'))
const dateToConfirm = f.format(selectedDate)

const formData = reactive({
    subject: "",
    email: localStoredUserData.email,
    phone: "",
    dni: "",
    hin: ""
})

watch(patientSelected, (loadedPatient) => {
    if (loadedPatient) {
        formData.phone = loadedPatient.patient.phone_number
        formData.dni = loadedPatient.patient.dni
        formData.hin = loadedPatient.patient.health_insurance
    }
})

const isFormValid = computed(() => {
    return (
        formData.subject &&
        formData.email &&
        formData.phone &&
        formData.dni &&
        formData.hin
    )
})

const submitForm = async () => {
    const year = selectedDate.getUTCFullYear();
    const month = String(selectedDate.getUTCMonth() + 1).padStart(2, '0');
    const day = String(selectedDate.getUTCDate()).padStart(2, '0');
    const hours = String(selectedDate.getUTCHours() - 3);
    const minutes = String(selectedDate.getUTCMinutes()).padStart(2, '0');
    const seconds = String(selectedDate.getUTCSeconds()).padStart(2, '0');
    // format: 'YYYY-MM-DDTHH:MM:SSZ'
    const formattedDate = `${year}-${month}-${day}T${hours}:${minutes}:${seconds}Z`;

    console.log(formattedDate);
    if (isFormValid.value) {
        console.log('Form submitted:', formData)
        const result = await saveAppointment(formattedDate, formData.subject, patientSelected.value.patient.id, localStoredUserData.email)
        if (result.error) {
            if (result.code === 401) {
                toast.info(result.message)
                router.replace({
                    path: '/registro',
                    query: { mode: 'login' }
                })
                return
            }
            console.error(result.message)
            toast.error('No se pudo agendar el turno. Intente nuevamente')
        } else {
            toast.success('Turno agendado')
            router.replace('/perfil')
        }
    }
}

</script>

<template>
    <div class="form-container">
        <h2>Confirma tu turno</h2>
        <p>{{ String(dateToConfirm).charAt(0).toUpperCase() + String(dateToConfirm).slice(1) }}</p>
        <div id="patient-select-container">
            <select v-if="patientsList" name="patient" id="patient-selection" v-model="indexSelected">
                <option v-for="(data, index) in patientsList" :key="index" :value="index">
                    {{ data.patient.first_name + ' ' + data.patient.last_name }}
                </option>
            </select>
            <button @click="router.push('/perfil')" id="btn-add-patient">Añadir</button>
        </div>
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="subject">Asunto (máx. 50 caracteres)</label>
                <input type="text" id="subject" v-model="formData.subject" required maxlength="50" />
            </div>
            <div class="form-group">
                <label for="email">Correo</label>
                <input type="email" id="email" :value="formData.email" disabled />
            </div>
            <div class="form-group">
                <label for="phone">Número de teléfono</label>
                <input type="tel" id="phone" v-model="formData.phone" disabled />
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
    </div>
</template>

<style scoped>
* {
    font-family: poppins-regular;
}

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

button[type="submit"] {
    width: 90%;
    padding: 12px;
    margin-top: 2em;
    background-color: #3790D0;
    font-weight: 600;
    color: white;
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

button[type="submit"]:disabled {
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

#patient-select-container {
    display: flex;
    align-items: center;
    justify-content: center;
    /* background-color: #333; */
}

#btn-add-patient {
    width: 30%;
    height: 30px;
    background-color: #3790D0;
    color: #EEEEEE;
    border: 4px solid transparent;
    border-radius: 0.5em;
    cursor: pointer;

    &:hover {
        border: 4px solid #2176b3;
    }

}

@media (max-width: 330px) {
    #patient-select-container {
        flex-direction: column;

        button,
        select {
            margin-top: 1em;
            width: 90%;
        }
    }
}
</style>