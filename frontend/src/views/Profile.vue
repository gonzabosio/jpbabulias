<script setup>
import { computed, onBeforeMount, reactive, ref, watch, watchEffect } from 'vue';
import { deletePatientById, editPatientData, getPatientsDataByUserId, savePatient } from '../fetch/patient';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { logout } from '../fetch/user';
import LoadingSpinner from '../components/LoadingSpinner.vue';

const toast = useToast()
const router = useRouter()
const showLoadingSpinner = ref(false)
const localUserData = ref(null)
const dbPatients = ref(null)
const personToShow = ref({})
const addPatientForm = ref(false)
const patientAppts = ref([])
const mainPatientIdx = ref(0)
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
        toast.error('No se cargaron tus datos correctamente. Recargue la página')
    }
    dbPatients.value = result.patients
    personToShow.value = dbPatients.value.find(it => it.patient.main)
    mainPatientIdx.value = dbPatients.value.findIndex(it => it.patient.main);
    patientAppts.value = personToShow.value.appointments
}

onBeforeMount(() => {
    const userData = localStorage.getItem('user')
    if (userData) {
        localUserData.value = JSON.parse(userData)
    }
})
watchEffect(async () => {
    if (localUserData.value) {
        await loadUserAndPatients()
    }
})

const selectedIdx = ref(0)
watch(selectedIdx, (currentIdx) => {
    personToShow.value = dbPatients.value[currentIdx]
    patientAppts.value = personToShow.value.appointments
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

const addPatient = async () => {
    if (isFormValid.value) {
        showLoadingSpinner.value = true
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
            showLoadingSpinner.value = false
            toast.error('No se pudo cargar el paciente. Intente nuevamente')
            return
        }
        await loadUserAndPatients()
        addPatientForm.value = false
        showLoadingSpinner.value = false
        toast.success('Paciente agregado')
    }
}

const handleLogout = async () => {
    showLoadingSpinner.value = true
    const result = await logout()
    if (result.error) {
        router.replace('/')
        return
    }
    localStorage.removeItem('user')
    router.replace('/')
}

const modifPatientForm = ref(false)
const handleModifPatientForm = (patient) => {
    formData.firstName = patient.first_name
    formData.lastName = patient.last_name
    formData.dni = patient.dni
    formData.phone = patient.phone_number.replace(/^\+54/, "")
    formData.hin = patient.health_insurance
    modifPatientForm.value = true
}

const submitModifPatientForm = async () => {
    showLoadingSpinner.value = true
    const result = await editPatientData(formData, personToShow.value.patient.id)
    if (result.error) {
        if (result.code === 401) {
            toast.info(result.message)
            router.replace({
                path: '/registro',
                query: { mode: 'login' }
            })
            return
        }
        showLoadingSpinner.value = false
        toast.error('No se pudo actualizar los datos del paciente')
    }
    modifPatientForm.value = false
    showLoadingSpinner.value = false
    toast.success('Datos del paciente actualizados')
}

const handleDeletePatient = async (patientId) => {
    showLoadingSpinner.value = true
    const result = await deletePatientById(patientId)
    if (result.error) {
        if (result.code === 401) {
            toast.info(result.message)
            router.replace({
                path: '/registro',
                query: { mode: 'login' }
            })
        } else if (result.code === 409) {
            showLoadingSpinner.value = false
            toast.error('El paciente tiene citas pendientes. No se puede eliminar')
        } else {
            showLoadingSpinner.value = false
            toast.error('No se pudo eliminar el paciente. Intente de nuevo')
        }
    }
    dbPatients.value = [...dbPatients.value.filter(patient => patient.patient.id !== patientId)]
    console.log(selectedIdx.value, mainPatientIdx.value)
    selectedIdx.value = mainPatientIdx.value
    showLoadingSpinner.value = false
    toast.success('Paciente eliminado')
}

const handleAddPatientFormSwitch = async () => {
    formData.firstName = ""
    formData.lastName = ""
    formData.phone = ""
    formData.dni = ""
    formData.hin = ""
    addPatientForm.value = !addPatientForm.value
}
</script>

<template>
    <div id="container" v-if="!showLoadingSpinner">
        <h1>Perfil</h1>
        <section id="user-data">
            <article>{{ localUserData.email }}</article>
        </section>

        <div v-if="dbPatients && localUserData && !addPatientForm && !modifPatientForm">
            <section id="patients-section">
                <button @click="handleAddPatientFormSwitch" id="btn-add-patient">&#43; Añadir paciente</button>
                <p>Seleccionar paciente</p>
                <select name="patients" id="select-patient" v-model="selectedIdx">
                    <option :value="index" v-for="(it, index) in dbPatients" :key="index">
                        {{ it.patient.first_name + ' ' + it.patient.last_name }}
                    </option>
                </select>
                <div id="patient-data" v-if="personToShow">
                    <ul>
                        <li>Teléfono: {{ personToShow.patient.phone_number }}</li>
                        <li>DNI: {{ personToShow.patient.dni }}</li>
                        <li>Obra Social: {{ personToShow.patient.health_insurance }}</li>
                    </ul>
                    <button @click="handleModifPatientForm(personToShow.patient)" class="btn-patient-mod">Editar
                        información</button>
                    <button @click="handleDeletePatient(personToShow.patient.id)" class="btn-patient-mod"
                        id="btn-del-patient" v-if="!personToShow.patient.main">Borrar
                        paciente</button>
                    <p v-if="patientAppts">Turnos:</p>
                    <div v-if="patientAppts" v-for="it in patientAppts" class="appt-card">
                        <span>{{ new Date(it.appt_date).toLocaleString('es-ar', {
                            timeZone: 'UTC', dateStyle: "short",
                            timeStyle: "short"
                        }) }}</span>
                    </div>
                </div>
                <p>❗Para cancelar un turno. Comuníquese directamente con el consultorio.</p>
            </section>
        </div>
        <div v-if="addPatientForm" class="form-container">
            <form @submit.prevent="addPatient">
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
            <button @click="handleAddPatientFormSwitch" id="btn-cancel">Cancelar</button>
        </div>
        <button v-if="!modifPatientForm && !addPatientForm" @click="handleLogout" id="btn-logout">Cerrar sesión</button>
        <div v-if="modifPatientForm" class="form-container">
            <form @submit.prevent="submitModifPatientForm">
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
                <button type="submit" :disabled="!isFormValid">Confirmar</button>
            </form>
            <button @click="modifPatientForm = false" id="btn-cancel">Cancelar</button>
        </div>
    </div>
    <div v-else class="loading-spinner-overlay">
        <LoadingSpinner />
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

#patients-section {
    margin: 1em 0;
    padding: 1em 2em;
    border-radius: 0.5em;
    border: 2px solid #c9c9c9;
}

#select-patient {
    width: 180px;
    border-radius: 0.5em;
    font-size: 1rem;
    height: 40px;
}

#patient-data {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    text-align: left;
    margin: 1em;
}

.appt-card {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 180px;
    padding: 0.5em;
    background-color: #fff;
    border: 2px solid gray;
    margin-bottom: 1em;
    border-radius: 0.5em;
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

.btn-patient-mod {
    cursor: pointer;
    border-radius: 0.5em;
    width: 200px;
    height: auto;
    padding: 0.5em;
    border: 4px solid transparent;
    background-color: #3790D0;
    transition: 0.25s;
    color: #EEEEEE;
    margin-bottom: 1em;
    font-weight: 600;
    font-size: 1rem;

    &:hover {
        border: 4px solid #2176b3;
    }
}

#btn-del-patient {
    background-color: transparent;
    border: 4px solid #3790D0;
    color: #3790D0;

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