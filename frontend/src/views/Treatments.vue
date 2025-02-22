<script setup>
import { ref } from "vue";
import { useRouter } from 'vue-router'
const router = useRouter()
const mediaZirconio = [
    { type: "image", src: new URL("../assets/treatments/zirconio/zirconio_1.jpeg", import.meta.url).href },
    { type: "image", src: new URL("../assets/treatments/zirconio/zirconio_2.jpg", import.meta.url).href },
    { type: "image", src: new URL("../assets/treatments/zirconio/zirconio_3.jpg", import.meta.url).href },
    { type: "video", src: new URL("../assets/treatments/zirconio/zirconio_vid.mp4", import.meta.url).href }
]
const mediaCarillas = [
    { type: "image", src: new URL("../assets/treatments/carillas/carillas_1.jpeg", import.meta.url).href },
    { type: "image", src: new URL("../assets/treatments/carillas/carillas_2.jpeg", import.meta.url).href },
    { type: "video", src: new URL("../assets/treatments/carillas/carillas_vid.mp4", import.meta.url).href }
]

const mediaOrtodoncia = [
    { type: "image", src: new URL("../assets/treatments/ortodoncia/ortodoncia_1.jpg", import.meta.url).href },
    { type: "image", src: new URL("../assets/treatments/ortodoncia/ortodoncia_2.jpg", import.meta.url).href },
    { type: "image", src: new URL("../assets/treatments/ortodoncia/ortodoncia_3.jpg", import.meta.url).href },
    { type: "image", src: new URL("../assets/treatments/ortodoncia/ortodoncia_4.jpg", import.meta.url).href },
    { type: "image", src: new URL("../assets/treatments/ortodoncia/ortodoncia_5.jpg", import.meta.url).href },
    { type: "image", src: new URL("../assets/treatments/ortodoncia/ortodoncia_6.jpg", import.meta.url).href }
]
const medias = ref([
    {
        name: 'Ortodoncia',
        currentIndex: 0,
        content: mediaOrtodoncia
    },
    {
        name: 'Carillas',
        currentIndex: 0,
        content: mediaCarillas
    },
    {
        name: 'Rehabilitación en zirconio',
        currentIndex: 0,
        content: mediaZirconio,
    },
])


const prev = (groupIndex) => {
    medias.value[groupIndex].currentIndex =
        (medias.value[groupIndex].currentIndex - 1 + medias.value[groupIndex].content.length) %
        medias.value[groupIndex].content.length
}

const next = (groupIndex) => {
    medias.value[groupIndex].currentIndex =
        (medias.value[groupIndex].currentIndex + 1) % medias.value[groupIndex].content.length
}
</script>

<template>
    <div id="container">
        <h1>Tratamientos</h1>
        <div class="carousel-container">
            <div class="carousel-group" v-for="(media, groupIndex) in medias" :key="groupIndex">
                <h2>{{ media.name }}</h2>
                <div class="carousel" :style="{ transform: `translateX(-${media.currentIndex * 100}%)` }">
                    <div v-for="(item, index) in media.content" :key="index" class="carousel-item">
                        <img v-if="item.type === 'image'" :src="item.src" alt="tratamiento" />
                        <video v-else autoplay loop muted playsinline>
                            <source :src="item.src" type="video/mp4" />
                        </video>
                    </div>
                </div>
                <button @click="prev(groupIndex)" class="prev">&#10094;</button>
                <button @click="next(groupIndex)" class="next">&#10095;</button>
            </div>
        </div>
        <div id="over-btn-appt">
            <button id="btn-appt" @click="router.push('/turnos')">Pedir turno</button>
        </div>
    </div>
</template>

<style scoped>
#container {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 7em 0;
    background-image: linear-gradient(to top, #e6e9f0 0%, #eef1f5 100%);
}

h1 {
    border: 2px solid #121215;
    padding: 1em;
    border-radius: 0.5em;
}

h2 {
    margin: 0;
    border-top: 2px solid #12121559;
    border-bottom: 2px solid #12121559;
    background-color: #f0f0f0;
    padding: 0.5em 0;
}

.carousel-container {
    width: 400px;
    margin: auto;
    overflow: hidden;
    box-sizing: border-box;
}

.carousel-group {
    margin-bottom: 80px;
    background-color: #f0f0f0;
    border-radius: 0.5em;
    border-bottom: 2px solid #12121559;

}

.carousel {
    display: flex;
    transition: transform 0.5s ease-in-out;
    background-color: rgb(211, 211, 211);
    border-bottom: 2px solid #12121559;
}

.carousel-item {
    display: flex;
    align-items: center;
    justify-content: center;
    min-width: 100%;
    box-sizing: border-box;

}

img,
video {
    width: 90%;
    height: 80%;
    display: block;
    margin: 8px;
    object-fit: contain;
}

button {
    background: #3790D0;
    color: white;
    border: none;
    border-radius: 0.5em;
    cursor: pointer;
    padding: 10px;
    font-size: 20px;
    margin: 5px;
    justify-content: space-between;
    transition: 0.25s;

    &:hover {
        background-color: #2176b3;
    }
}

#over-btn-appt {
    padding: 1em;
    border-radius: 0.5em;
    border: 2px solid #3790D0;
}

#btn-appt {
    font-family: poppins-regular;
    width: 300px;
    height: auto;
    font-weight: 600;
}

@media (max-width: 450px) {
    .carousel-container {
        width: 300px;
    }

    #btn-appt {
        width: 200px;
    }
}

@media (max-width: 350px) {
    h1 {
        border: 2px solid #121215;
        padding: 0.5em;
        border-radius: 0.5em;
        font-size: 1.5rem;
    }

    #over-btn-appt {
        padding: 0.5em;
    }

    h2 {
        font-size: 1rem;
    }

    .carousel-container {
        overflow: hidden;
        width: 250px;
    }
}
</style>