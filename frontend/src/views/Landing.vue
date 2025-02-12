<script setup>
import Footer from '../components/Footer.vue'

const descriptionList = [
    'Más de XX años de experiencia como profesional',
    'Especializado en cierta práctica',
    'Recibido en la Universidad de Córdoba'
]
</script>

<template>
    <div>
        <div id="leading-container">
            <h1 class="shiny-title">ODONTOLOGÍA DIGITAL</h1>
            <p>Implantes · Ortodoncia · Zirconio · Modelado 3D</p>
            <div id="navigation-container">
                <RouterLink to="/turnos" class="nav-option">Pedir turno</RouterLink>
                <RouterLink to="/tratamientos" class="nav-option">Tratamientos</RouterLink>
            </div>
            <section id="about">
                <article id="dentist-pic">
                    <img src="../assets/dentist_pic_dec.png" alt="dentista">
                </article>
                <article id="description">
                    <ul>
                        <li class="description-items" v-for="item in descriptionList">{{ item }}</li>
                    </ul>
                </article>
            </section>
        </div>
    </div>
</template>
<style scoped>
#leading-container {
    margin: 8px;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.shiny-title {
    font-family: poppins-bold;
    font-size: 4rem;
}

p {
    font-size: 1.5rem;
    color: #404042;
}

#navigation-container {
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    margin: 4em 0 4em 0;
}

.nav-option {
    text-decoration: none;
    color: #EEEEEE;
    padding: 0.7em;
    font-size: 1.7rem;
    border-radius: 0.8em;
    background-color: #3790D0;
    margin: 8px;
    width: 250px;
    transition: 0.25s;
    will-change: transform;
    border: 4px solid #3790D0;

    &:hover {
        transform: scale(1.05);
    }

    &:last-child {
        background-color: transparent;
        color: #3790D0;
    }
}

.shiny-title {
    position: relative;
    color: #ffffff;
    background: linear-gradient(to right, #2d2d30 0%, white 20%, #121215 100%);
    background-size: 200% auto;
    background-position: -100% center;
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    text-decoration: none;
    white-space: nowrap;
    animation: shine 3s infinite linear;
}

@keyframes shine {
    0% {
        background-position: -100% center;
    }

    100% {
        background-position: 100% center;
    }
}

#about {
    display: flex;
    align-items: center;
    padding: 1.8em;
    margin-top: 6em;
    margin-bottom: 6em;
    border-left: 4px solid #3790D0;
    border-right: 4px solid #3790D0;
    mask-image: linear-gradient(to bottom, transparent 0%, white 10%, white 90%, transparent 100%);

    #description {
        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 32px;
    }

    .description-items {
        text-align: left;
        font-size: 1.5rem;
        padding: 16px 16px;
        border-radius: 8px;
        transition: transform 0.3s ease, box-shadow 0.3s ease;
    }

    .description-items:hover {
        transform: scale(1.05);
        box-shadow: 0 0 10px #3790d07a;
    }

    #dentist-pic {
        display: flex;
        margin-right: 16px;

        img {
            width: 250px;
        }
    }
}


@media (max-width: 800px) {
    .shiny-title {
        font-size: 3rem;
    }

    p {
        font-size: 1.2rem;
    }

    .nav-option {
        padding: 0.8em;
        width: 170px;
        font-size: 20px;
    }

    #about {
        .description-items {
            font-size: 1.2rem;
        }
    }

}

@media (max-width: 650px) {
    .shiny-title {
        font-size: 2.2rem;
    }

    p {
        font-size: 0.9rem;
    }

    .nav-option {
        padding: 0.6em;
        width: 160px;
        font-size: 20px;
    }

    #about {
        flex-direction: column;

        .description-items {
            font-size: 1rem;
        }

        #dentist-pic img {
            width: 250px;
        }
    }


}

@media (max-width: 500px) {
    .shiny-title {
        font-size: 1.6rem;
    }

    .nav-option {
        padding: 0.6em;
        width: 150px;
        font-size: 1rem;
    }
}

@media (max-width: 350px) {
    .shiny-title {
        font-size: 1.1rem;
        text-wrap: wrap;
    }

    .nav-option {
        padding: 0.5em;
        width: 120px;
        font-size: 16px;
    }

    #about {
        .description-items {
            font-size: 0.8rem;
        }

        #dentist-pic img {
            width: 150px;
        }
    }
}
</style>