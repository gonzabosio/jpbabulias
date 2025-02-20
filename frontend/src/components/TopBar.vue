<script setup>
import { onBeforeMount, ref, watch } from 'vue';
import { isLoggedIn, userData, switchAuthStatus } from '../components/scripts/topBarRef'
onBeforeMount(() => {
    switchAuthStatus()
})

const profileLett = ref('')
watch(userData, () => {
    if (userData.value) {
        profileLett.value = String(userData.value.email).at(0).toLocaleUpperCase()
    }
})
</script>

<template>
    <header>
        <a href="/">
            <img src="../assets/tooth_icon_title.svg" alt="diente" id="tooth">
        </a>
        <div id="doctor-name">
            Dr. Juan Pablo Babulias
        </div>
        <div v-if="!isLoggedIn" id="register-container">
            <RouterLink :to="{ path: '/registro', query: { mode: 'login' } }" id="nav-login" class="underline-anim">
                Iniciar Sesión
            </RouterLink>
            <RouterLink :to="{ path: '/registro', query: { mode: 'signup' } }" id="nav-signup" class="underline-anim">
                Registrarse
            </RouterLink>
        </div>
        <div v-else class="header-right">
            <RouterLink to="/perfil" id="nav-profile">
                <span id="letter">{{ profileLett }}</span>
            </RouterLink>
        </div>
    </header>
</template>

<style scoped>
header {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    background-color: #d9dee9;
    color: #121215;
    display: flex;
    align-items: center;
    justify-content: space-between;
    text-align: left;
    border-bottom: 1px solid rgb(150, 150, 150);
    z-index: 100;
}

a {
    display: flex;
    align-items: center;
    justify-content: center;

    img {
        transition: 0.2s;
    }

    &:hover {
        img {
            filter: drop-shadow(4px 4px 6px rgba(0, 0, 0, 0.5));
        }
    }
}

#tooth {
    width: 50px;
    margin: 8px 16px 8px 16px;
}

#doctor-name {
    font-size: 2.5rem;
    font-weight: bold;
    flex-grow: 1;
}

.header-right {
    display: flex;
}

#nav-profile {
    display: flex;
    cursor: pointer;
    border-radius: 100%;
    border: 2px solid #121215;
    margin: 8px 16px 8px 8px;
    transition: box-shadow 0.2s ease-in-out;
    transition: 0.2s;
    text-decoration: none;
    color: #121215;
    width: 50px;
    height: 50px;
    text-align: center;
    align-items: center;
    justify-content: center;

    &:hover {
        box-shadow: 0 0 0 2px #121215;
        background-color: #EEEEEE;
    }

    span {
        font-size: 2rem;
    }
}

#register-container {
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 16px;
    padding: 1em;
}

.underline-anim {
    display: inline-block;
    position: relative;
    text-decoration: none;
    color: #121215;
    font-size: 18px;
}

.underline-anim::after {
    content: '';
    position: absolute;
    width: 100%;
    transform: scaleX(0);
    height: 2px;
    bottom: 0;
    left: 0;
    background-color: #121215;
    transition: transform 0.25s ease-out;
}

.underline-anim:hover::after {
    transform: scaleX(1);
}

.underline-anim::after {
    transform-origin: bottom center;
}

.underline-anim:hover::after {
    transform-origin: bottom center;
}

hr {
    margin: 0;
}

@media (max-width: 650px) {
    #doctor-name {
        font-size: 28px;
    }
}

@media (max-width: 400px) {
    #tooth {
        width: 30px;
    }

    #doctor-name {
        font-size: 20px;
    }

    #nav-profile {
        padding: 0;
    }

    #register-container {
        flex-direction: column;
    }

    .underline-anim {
        font-size: 16px;
    }
}
</style>