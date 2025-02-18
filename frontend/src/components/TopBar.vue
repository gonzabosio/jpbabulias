<script setup>
const isLoggedIn = true
</script>

<template>
    <header>
        <a href="/">
            <img src="../assets/tooth_icon_title.svg" alt="diente" id="tooth">
        </a>
        <div id="doctor-name">
            Dr. Juan Pablo Babulias
        </div>
        <div v-if="isLoggedIn" id="register-container">
            <RouterLink :to="{ path: '/registro', query: { mode: 'login' } }" id="nav-login" class="underline-anim">
                Iniciar Sesión
            </RouterLink>
            <RouterLink :to="{ path: '/registro', query: { mode: 'signup' } }" id="nav-signup" class="underline-anim">
                Registrarse
            </RouterLink>
        </div>
        <div v-else class="header-right">
            <RouterLink to="/perfil" id="nav-profile">
                <svg xmlns="http://www.w3.org/2000/svg" width="45" height="45" viewBox="0 0 24 24"
                    style="fill: rgba(18, 18, 21, 1)">
                    <path
                        d="M12 2a5 5 0 1 0 5 5 5 5 0 0 0-5-5zm0 8a3 3 0 1 1 3-3 3 3 0 0 1-3 3zm9 11v-1a7 7 0 0 0-7-7h-4a7 7 0 0 0-7 7v1h2v-1a5 5 0 0 1 5-5h4a5 5 0 0 1 5 5v1z">
                    </path>
                </svg>
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
    border-radius: 50%;
    border: 2px solid #121215;
    padding: 0.3em;
    margin: 8px 16px 8px 8px;
    transition: box-shadow 0.2s ease-in-out;
    transition: 0.2s;
    text-decoration: none;
    color: #121215;
    /* width: 50px;
    height: 50px;
    text-align: center;
    align-items: center;
    justify-content: center; */

    &:hover {
        box-shadow: 0 0 0 2px #121215;
        background-color: #EEEEEE;
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