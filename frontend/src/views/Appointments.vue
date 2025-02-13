<script setup>
import { ref } from 'vue'
import { useToast } from "vue-toastification";
import { useRouter } from 'vue-router'

const router = useRouter()
const toast = useToast()
const selectedDate = ref(new Date())

const disabledDates = ref([
    {
        repeat: {
            weekdays: [1, 7]
        },

    }
])

const attrs = ref([
    {
        key: 'occupied',
        highlight: {
            color: 'red',
        },
        dates: [new Date(2025, 2, 14)],
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

const onDayClick = (day) => {
    hourSelected.value = ''
    // console.log('Day ID: ', day.id)
    const clickedDate = new Date(day.date)
    console.log('Clicked date:', clickedDate)
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
                console.log('Date is fully occupied 🔒')
                return
            }
        } else if (typeof attr.dates === 'object' && attr.dates.start && attr.dates.end) {
            // if clicked day is within the date range
            isWithinRange = clickedDate >= attr.dates.start && clickedDate <= attr.dates.end
            if (isWithinRange) {
                console.log('Closed for holidays 🌴')
                return
            }
        }
    })
    if (!isDateInArray && !isWithinRange) {
        console.log('send request for the day: ', selectedDate.value.toISOString())
        // const formattedDate = selectedDate.value.toISOString().slice(0, 19).replace("T", " ")
    }
}

const normalizeDate = (date) => {
    if (!date) return null;
    const d = new Date(date);
    d.setHours(0, 0, 0, 0);
    return d.getTime();
}

const hoursList = ['07:00', '07:30', '08:00']
const hourSelected = ref('')

const doSomething = (selDate) => {
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
            <p>Seleccione una hora</p>
            <div id="appt-list">
                <div class="appt-card" v-for="(hour, index) in hoursList" :key="index" @click="hourSelected = hour">
                    <input type="radio" class="btn-radio" :id="'appt-' + index" name="appointment" :value="hour"
                        v-model="hourSelected">
                    <label :for="'appt-' + index">{{ hour }}</label>
                </div>
                <button id="btn-schedule" @click="doSomething(selectedDate)"
                    :disabled="hourSelected === ''">Agendar</button>
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
    display: flex;
    flex-direction: column;
    align-items: center;
}

.appt-card {
    cursor: pointer;
    display: flex;
    align-items: center;
    border: 2px solid rgb(175, 175, 175);
    border-radius: 0.5em;
    padding: 4px 8px;
    margin: 4px 0px;

    input[type=radio] {
        cursor: pointer;
        width: 28px;
        height: 28px;
    }

    label {
        cursor: pointer;
    }
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
    width: 100px;
    height: 50px;
    border-radius: 0.5em;
    text-decoration: none;
    display: flex;
    align-items: center;
    justify-content: center;

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