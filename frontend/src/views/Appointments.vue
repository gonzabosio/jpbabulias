<script setup>
import { computed, onBeforeMount, onMounted, ref } from 'vue'
import { useToast } from "vue-toastification";
import { useRouter } from 'vue-router'
import ThreeDotsMove from '../components/ThreeDotsMove.vue';
import { getDayAppointments, getFullyBookedDates } from '../fetch/appt';

const router = useRouter()
const toast = useToast()
const fullDates = ref([])
onBeforeMount(async () => {
    // const result = await getFullyBookedDates()
    // console.log(result.fully_booked_dates)
    // fullDates.value = result.fully_booked_dates.map(date => new Date(date))
})

const selectedDate = ref(new Date())
const disabledDates = ref([
    {
        repeat: {
            weekdays: [1, 7]
        },

    }
])

const attrs = computed(() => [
    {
        key: 'fullDates',
        highlight: {
            color: 'red',
        },
        dates: fullDates.value,
    },
    {
        key: 'holidays',
        highlight: {
            color: 'orange'
        },
        dates: { start: new Date(2025, 1, 10), end: new Date(2025, 1, 23) },
    }
])

const minDate = new Date()
minDate.setDate(minDate.getDate() + 1)
const maxDate = new Date()
maxDate.setMonth(maxDate.getMonth() + 4)

const hoursList = ref([])
const hourSelected = ref('')
let isLoading = ref(false)

const monToThuHours = ['08:00', '08:30', '09:00', '09:30', '10:00', '10:30', '11:00', '11:30', '16:00', '16:30', '17:00', '17:30', '18:00', '18:30', '19:00', '19:30']
const fridayHours = ['08:00', '08:30', '09:00', '09:30', '10:00', '10:30', '11:00', '11:30', '12:00', '12:30', '13:00', '13:30', '14:00', '14:30', '15:00', '15:30']

const onDayClick = async (day) => {
    hourSelected.value = ''
    // console.log('Day ID: ', day.id)
    const clickedDate = new Date(day.date)
    // console.log('Clicked date:', clickedDate)
    selectedDate.value = clickedDate

    if (day.isDisabled) {

        console.log('Date is unavailable ❌')
        return
    }
    let isDateInArray = false
    let isWithinRange = false
    attrs.value.forEach(attr => {
        if (Array.isArray(attr.dates)) {
            // if clicked day is in the array of dates
            isDateInArray = attr.dates.some(date => date.toISOString().split('T')[0] === day.date.toISOString().split('T')[0])
            if (isDateInArray) {
                toast.warning('La fecha seleccionada esta completamente llena')
                // console.log('Date is fully booked 🔒')
                return
            }
        } else if (typeof attr.dates === 'object' && attr.dates.start && attr.dates.end) {
            // if clicked day is within the date range
            isWithinRange = clickedDate >= attr.dates.start && clickedDate <= attr.dates.end
            if (isWithinRange) {
                toast.warning('El consultorio no abre durante el día seleccionado')
                // console.log('Closed for holidays 🌴')
                return
            }
        }
    })
    if (!isDateInArray && !isWithinRange) {
        isLoading.value = true
        const result = await getDayAppointments(selectedDate.value.toISOString())
        if (result.error) {
            console.error('Failed to get available appointments', result.message)
            isLoading.value = false
            toast.error('Ocurrió un error al intentar cargar los horarios', {
                timeout: 7000
            })
        } else {
            setTimeout(() => {
                let takenAppts = result.appointments
                // console.log(takenAppts)
                const dayNum = selectedDate.value.getDay()
                if (takenAppts.length === 0) {
                    if (!isFriday(dayNum)) {
                        hoursList.value = monToThuHours
                    } else {
                        hoursList.value = fridayHours
                    }
                } else {
                    takenAppts = result.appointments
                    hoursList.value = !isFriday(dayNum)
                        ? monToThuHours.filter(hour => !takenAppts.includes(hour))
                        : hoursList.value = fridayHours.filter(hour => !takenAppts.includes(hour))

                }
                // console.log(hoursList.value)
                isLoading.value = false
            }, 500)
        }
    }
}

const isFriday = (day) => {
    if (day > 0 && day < 5) {
        // console.log('lunes-jueves')
        return false
    } else {
        // console.log('viernes')
        return true
    }
}

const normalizeDate = (date) => {
    if (!date) return null;
    const d = new Date(date);
    d.setHours(0, 0, 0, 0);
    return d.getTime();
}

const goConfirm = (selDate) => {
    let formattedDate = new Date(selDate)
    const [hour, minute] = hourSelected.value.split(':')
    formattedDate.setHours(hour, minute)
    sessionStorage.setItem('selectedDate', formattedDate)
    router.push('/turnos/confirmar')
}
</script>

<template>
    <div id="container">
        <h2>Consulte los horarios disponibles</h2>
        <div id="references">
            <div class="reference">
                <span id="gray" class="circle"></span>No hábiles
            </div>
            <div class="reference">
                <span id="red" class="circle"></span>Sin turnos disponibles
            </div>
            <div class="reference">
                <span id="orange" class="circle"></span>Consultorio cerrado
            </div>
        </div>
        <VDatePicker id="appt-calendar" v-model.string="selectedDate" transparent :is-dark="false" locale="es"
            :disabled-dates="disabledDates" expanded :rows="2" :first-day-of-week="0" :min-date="minDate"
            :attributes="attrs" @dayclick="onDayClick" />
        <div v-if="normalizeDate(selectedDate) >= normalizeDate(minDate)" id="schedule">
            <p>
                {{
                    selectedDate.getUTCDate() + '/' + (selectedDate.getUTCMonth() + 1) + '/' + selectedDate.getUTCFullYear()
                }}
            </p>
            <div v-if="isLoading" id="load">
                <ThreeDotsMove />
            </div>
            <div v-else-if="hoursList.length !== 0">
                <p>Seleccione una hora</p>
                <div id="appt-list">
                    <div class="appt-card" v-for="(hour, index) in hoursList" :key="index" @click="hourSelected = hour">
                        <input type="radio" class="btn-radio" :id="'appt-' + index" name="appointment" :value="hour"
                            v-model="hourSelected">
                        <label :for="'appt-' + index">{{ hour }}</label>
                    </div>
                    <button id="btn-schedule" @click="goConfirm(selectedDate)"
                        :disabled="hourSelected === ''">Agendar</button>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
#container {
    margin-top: 12px;
}

.reference {
    display: flex;
    align-items: center;
    text-align: left;
    margin: 1em;
}

.circle {
    width: 24px;
    height: 24px;
    border-radius: 50%;
    margin-right: 8px;
}

#orange {
    background-color: #f97316;
}

#red {
    background-color: #dc2626;
}

#gray {
    background-color: #cbd5e1;
}

#appt-list {
    display: grid;
    align-items: center;
    justify-content: center;
    flex-wrap: wrap;
    gap: 10px;
    grid-template-columns: repeat(2, 0fr);
}

.appt-card {
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    border: 2px solid rgb(175, 175, 175);
    border-radius: 0.5em;
    padding: 4px 0px;
    width: 100px;

    input[type=radio] {
        cursor: pointer;
        width: 28px;
        height: 28px;
    }

    label {
        cursor: pointer;
    }
}

#btn-schedule-container {
    display: flex;
    width: 100%;
    background-color: #2176b3;
}

#btn-schedule {
    cursor: pointer;
    color: #EEEEEE;
    background-color: #3790D0;
    border: 4px solid #3790D0;
    transition: 0.25s;
    font-size: 1rem;
    margin: 1.5em 0;
    font-weight: 600;
    width: 100%;
    height: 50px;
    border-radius: 0.5em;
    grid-column: span 2;
    justify-self: center;

    &:disabled {
        background-color: gray;
        border: 4px solid gray;

        &:hover {
            border: 4px solid gray;
        }
    }

    &:hover {
        border: 4px solid #2176b3;
    }
}

.disabled-link {
    pointer-events: none;
    opacity: 0.5;
}


@media (max-width: 300px) {
    #container {
        padding: 8px;
    }

    h2 {
        font-size: 1.2rem;
    }

    .reference {
        font-size: 0.8rem;
    }

}
</style>