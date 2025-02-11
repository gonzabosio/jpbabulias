<script setup></script>

<template>
    <header>
        <img src="../assets/tooth_icon_title.svg" alt="diente" id="tooth">
        <div id="doctor-name">
            Dr. Juan Pablo Babulias
        </div>
        <div class="header-right">
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
    background-color: #d9dee9;
    color: #121215;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1em;
    text-align: left;
    border-bottom: 1px solid rgb(150, 150, 150);
}

#tooth {
    width: 50px;
}

#doctor-name {
    font-size: 2.5rem;
    font-weight: bold;
    margin-left: 16px;
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
    transition: box-shadow 0.2s ease-in-out;
    transition: 0.2s;
    text-decoration: none;

    &:hover {
        box-shadow: 0 0 0 2px #121215;
        background-color: #EEEEEE;
    }
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

    #nav-profile {
        padding: 0;
    }
}

@media (max-width: 315px) {
    #doctor-name {
        font-size: 20px;
    }
}
</style>