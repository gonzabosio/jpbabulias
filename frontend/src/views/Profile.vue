<script setup>
import { computed, onBeforeMount, onMounted, reactive, ref, watch } from 'vue';
import { getPatientsDataByUserId, savePatient } from '../fetch/patient';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';

const toast = useToast()
const router = useRouter()

const localUserData = ref(null)
const dbPatients = ref(null)
const patientToShow = ref({})
const addPatientForm = ref(false)

const loadUserAndPatients = async () => {
    const result = await getPatientsDataByUserId(localUserData.value["user_id"])
    if (result.error) {
        if (result.code === 401) {
            toast.info(result.message)
            router.replace({
                path: '/registro',
                query: { mode: 'login' }
            })
            return
        }
        console.log(result.message)
        toast.error('No se cargaron tus datos correctamente. Recargue la página')
    }
    // console.log(result.patients)
    dbPatients.value = result.patients
    patientToShow.value = dbPatients.value.find(patient => patient.main)
    // console.log(patientToShow.value)
}

onBeforeMount(() => {
    localUserData.value = JSON.parse(localStorage.getItem('user'))
})
onMounted(() => {
    loadUserAndPatients()
})

const selectedIdx = ref(0)
watch(selectedIdx, (currentIdx) => {
    patientToShow.value = dbPatients.value[currentIdx]
})

const formData = reactive({
    firstName: "",
    lastName: "",
    phone: "",
    dni: "",
    hin: ""
})

const isFormValid = computed(() => {
    return (
        formData.firstName &&
        formData.lastName &&
        formData.phone &&
        formData.dni &&
        formData.hin
    )
})

const submitForm = async () => {
    if (isFormValid.value) {
        let cleanedPhoneNumber = formData.phone.replace(/[^\d]/g, '');
        formData.phone = cleanedPhoneNumber
        const result = await savePatient(formData, localUserData.value["user_id"])
        if (result.error) {
            if (result.code === 401) {
                toast.info(result.message)
                router.replace({
                    path: '/registro',
                    query: { mode: 'login' }
                })
                return
            }
            toast.error('No se pudo cargar el paciente. Intente nuevamente')
            return
        }
        await loadUserAndPatients()
        addPatientForm.value = false
        toast.success('Paciente agregado')
    }
}
</script>

<template>
    <div id="container">
        <h1>Perfil</h1>
        <section id="user-data">
            <article>{{ localUserData.email }}</article>
        </section>

        <div v-if="dbPatients && localUserData && !addPatientForm">
            <section id="patients-sections">
                <p>Seleccionar paciente</p>
                <select name="patients" id="select-patient" v-model="selectedIdx">
                    <option :value="index" v-for="(patient, index) in dbPatients" :key="index">
                        {{ patient.first_name + ' ' + patient.last_name }}
                    </option>
                </select>
                <div id="patient-data" v-if="patientToShow">
                    <ul>
                        <li>Teléfono: {{ patientToShow.phone_number }}</li>
                        <li>DNI: {{ patientToShow.dni }}</li>
                        <li>Obra Social: {{ patientToShow.health_insurance }}</li>
                    </ul>
                </div>
                <button @click="addPatientForm = true" id="btn-add-patient">Añadir paciente</button>
            </section>
        </div>
        <div v-if="addPatientForm" class="form-container">
            <form @submit.prevent="submitForm">
                <div class="form-group">
                    <label for="firstName">Nombre</label>
                    <input type="text" id="firstName" v-model="formData.firstName" required />
                </div>
                <div class="form-group">
                    <label for="lastName">Apellido</label>
                    <input type="text" id="lastName" v-model="formData.lastName" required />
                </div>
                <div class="form-group">
                    <label for="phone">Número de teléfono</label>
                    <div class="phone-input-container">
                        <span class="phone-prefix">+54</span>
                        <input type="tel" id="phone" v-model="formData.phone" required placeholder="3564101010"
                            maxlength="15" />
                    </div>
                </div>
                <div class="form-group">
                    <label for="dni">DNI</label>
                    <input type="number" id="dni" v-model="formData.dni" required />
                </div>
                <div class="form-group">
                    <label for="hin">Obra social</label>
                    <input type="text" id="hin" v-model="formData.hin" required />
                </div>
                <button type="submit" :disabled="!isFormValid">Añadir</button>
            </form>
            <button @click="addPatientForm = false" id="btn-cancel">Cancelar</button>
        </div>
        <button @click="router.replace('/')" id="btn-logout">Cerrar sesión</button>

    </div>
</template>

<style scoped>
* {
    font-family: poppins-regular;
}

#container {
    padding-top: 6em;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}

#btn-add-patient {
    width: 180px;
    margin-top: 1em;
    background-color: #3790D0;
    font-weight: 600;
    color: white;
    border: none;
    border-radius: 0.5em;
    font-size: 1rem;
    cursor: pointer;
    transition: 0.25s;

    border: 4px solid transparent;

    &:hover {
        border: 4px solid #2176b3;
    }
}

#patient-data {
    text-align: left;
}

.form-container {
    margin: 1em auto;
    padding: 30px;
    background-color: #f5f5f5;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
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

button[type="submit"] {
    width: 90%;
    margin-top: 1em;
    background-color: #3790D0;
    font-weight: 600;
    color: white;
    border: none;
    border-radius: 0.5em;
    font-size: 1rem;
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

#btn-cancel,
#btn-logout {
    cursor: pointer;
    border: 2px solid rgb(224, 26, 26);
    border-radius: 0.5em;
    font-size: 1rem;
    color: rgb(224, 26, 26);
    background-color: transparent;
    transition: 0.25s ease;
    margin: 1em 0;

    &:hover {
        background-color: rgba(224, 26, 26, 0.100);
    }
}

#btn-cancel {
    width: 90%;
}

.phone-input-container {
    display: flex;
    align-items: center;
    border: 1px solid #ddd;
    border-radius: 0.5em;
    overflow: hidden;
}

.phone-prefix {
    /* background-color: #f0f2f5; */
    color: #555;
    margin: 8px;
}

.phone-input-container input {
    border: none;
    flex-grow: 1;
}
</style>