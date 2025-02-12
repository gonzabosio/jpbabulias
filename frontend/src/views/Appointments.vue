<script setup>
import { ref } from 'vue'
import { useToast } from "vue-toastification";

const toast = useToast()
const selectedDate = ref(new Date())
const f = new Intl.DateTimeFormat('es-ar')

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
    // console.log('Day ID: ', day.id)
    const clickedDate = new Date(day.date)
    console.log('Clicked date:', clickedDate.toUTCString())
    selectedDate.value = clickedDate

    if (day.isDisabled) {
        console.log('Date is unavailable ❌')
        return
    }
    attrs.value.forEach(attr => {
        if (Array.isArray(attr.dates)) {
            // if clicked day is in the array of dates
            const isDateInArray = attr.dates.some(date => date.toISOString().split('T')[0] === day.date.toISOString().split('T')[0]);
            if (isDateInArray) {
                console.log('Date is fully occupied 🔒')
            }
        } else if (typeof attr.dates === 'object' && attr.dates.start && attr.dates.end) {
            // if clicked day is within the date range
            const isWithinRange = clickedDate >= attr.dates.start && clickedDate <= attr.dates.end;
            if (isWithinRange) {
                console.log('Closed for holidays 🌴')
            }
        }
    });
}

const normalizeDate = (date) => {
    if (!date) return null;
    const d = new Date(date);
    d.setHours(0, 0, 0, 0);
    return d.getTime();
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
        <!-- <div id="appt-calendar-container"> -->
        <VDatePicker id="appt-calendar" v-model.string="selectedDate" transparent :is-dark="false" locale="es"
            :disabled-dates="disabledDates" expanded :rows="2" :first-day-of-week="0" :min-date="minDate"
            :attributes="attrs" @dayclick="onDayClick" />
        <!-- </div> -->
        <p v-if="normalizeDate(selectedDate) >= normalizeDate(minDate)">
            {{
                selectedDate.getUTCDate() + '/' + (selectedDate.getUTCMonth() + 1) + '/' + selectedDate.getUTCFullYear()
            }}
        </p>
    </div>
</template>

<style scoped>
#container {
    padding: 24px;
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