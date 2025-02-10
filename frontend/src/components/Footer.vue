<template>
    <footer>
        <div id="contact">
            <span>Contacto</span>
            <a id="instagram" href="https://www.instagram.com/od.jpbabulias" target="_blank">
                <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none"
                    stroke="#121215" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-brand-instagram">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <path d="M4 8a4 4 0 0 1 4 -4h8a4 4 0 0 1 4 4v8a4 4 0 0 1 -4 4h-8a4 4 0 0 1 -4 -4z" />
                    <path d="M9 12a3 3 0 1 0 6 0a3 3 0 0 0 -6 0" />
                    <path d="M16.5 7.5v.01" />
                </svg>
                <span class="det">Instagram</span>
            </a>
            <a id="whatsapp" href="https://wa.me/543564590331" target="_blank">
                <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none"
                    stroke="#121215" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-brand-whatsapp">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <path d="M3 21l1.65 -3.8a9 9 0 1 1 3.4 2.9l-5.05 .9" />
                    <path
                        d="M9 10a.5 .5 0 0 0 1 0v-1a.5 .5 0 0 0 -1 0v1a5 5 0 0 0 5 5h1a.5 .5 0 0 0 0 -1h-1a.5 .5 0 0 0 0 1" />
                </svg>
                <span class="det">Whatsapp</span>
            </a>
        </div>
        <div id="location">
            <span>Ubicación</span>
            <div class="location-details">
                <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none"
                    stroke="#121215" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-map-pin">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <path d="M9 11a3 3 0 1 0 6 0a3 3 0 0 0 -6 0" />
                    <path d="M17.657 16.657l-4.243 4.243a2 2 0 0 1 -2.827 0l-4.244 -4.243a8 8 0 1 1 11.314 0z" />
                </svg>
                <span class="det">San Francisco, Córdoba, Argentina</span>
            </div>
            <div class="location-details">
                <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none"
                    stroke="#121215" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"
                    class="icon icon-tabler icons-tabler-outline icon-tabler-building">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <path d="M3 21l18 0" />
                    <path d="M9 8l1 0" />
                    <path d="M9 12l1 0" />
                    <path d="M9 16l1 0" />
                    <path d="M14 8l1 0" />
                    <path d="M14 12l1 0" />
                    <path d="M14 16l1 0" />
                    <path d="M5 21v-16a2 2 0 0 1 2 -2h10a2 2 0 0 1 2 2v16" />
                </svg>
                <span class="det">Juan de Garay 2055</span>
            </div>
        </div>
    </footer>
</template>

<style scoped>
footer {
    background-color: #D9DEE9;
    position: fixed;
    display: flex;
    bottom: 0px;
    left: 0px;
    right: 0px;
    margin-bottom: 0px;
    padding-bottom: 2rem;
    height: fit-content;
    padding-left: 30%;
    padding-right: 30%;
}

hr {
    margin: 0;
}

span {
    font-size: 1.5rem;
}

.det {
    font-size: 1.2rem;
    margin-left: 4px;
    text-align: start;
}

#contact {
    display: flex;
    flex-direction: column;
    margin-top: 2rem;
    gap: 10px;
    text-align: left;
    flex-grow: 1;

    #whatsapp,
    #instagram {
        display: flex;
        align-items: center;
        text-decoration: none;
        color: #121215;

        &:hover {
            text-decoration: underline;
        }
    }

}

#location {
    display: flex;
    flex-direction: column;
    align-items: start;
    margin-top: 2rem;
    gap: 10px;
    text-align: right;

    .location-details {
        display: flex;
        align-items: center;
        color: #121215;
    }

}

@media (max-width: 1360px) {
    footer {
        padding-left: 20%;
        padding-right: 20%;
    }
}

@media (max-width: 900px) {
    footer {
        padding-left: 10%;
        padding-right: 10%;
    }
}

@media (max-width: 650px) {
    span {
        font-size: 1rem;
    }

    .det {
        font-size: 0.7rem;
    }
}

@media (max-width: 450px) {
    footer {
        flex-direction: column;
    }
}

@media (max-width: 300px) {
    .det {
        justify-content: start;
    }
}
</style>
