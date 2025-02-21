<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { sendPrompt } from '../fetch/bot'

const isChatOpen = ref(false)
const messages = ref([])
const userInput = ref('')

const toggleChat = () => {
    isChatOpen.value = !isChatOpen.value
    if (isChatOpen.value) {
        nextTick(() => {
            scrollToBottom()
        })
    }
}

const chatContainer = ref(null)
const isProcessing = ref(false)
const scrollToBottom = () => {
    if (chatContainer.value) {
        chatContainer.value.scrollTop = chatContainer.value.scrollHeight;
    }
};

const sendMessage = () => {
    if (userInput.value.trim() === '') return

    messages.value.push({ text: userInput.value, sender: 'user' })
    nextTick(() => {
        scrollToBottom()
    })
    const prompt = userInput.value
    userInput.value = ''

    isProcessing.value = true
    setTimeout(async () => {
        const result = await sendPrompt(prompt)
        if (result.error) {
            console.error(result)
            messages.value.push({ text: 'Ocurrió un error. Intente de nuevo.', sender: 'bot' })
            isProcessing.value = false
        } else {
            // console.log(result)
            messages.value.push({ text: result.response, sender: 'bot' })
            isProcessing.value = false
            nextTick(() => {
                scrollToBottom()
            })
        }
    }, 500)
}

onMounted(() => {
    messages.value.push({ text: '¡Hola! ¿Cómo puedo ayudarte?', sender: 'bot' })
})
</script>

<template>
    <div class="chatbot-container">
        <button @click="toggleChat" class="chatbot-button" :class="{ 'open': isChatOpen }">
            <span v-if="!isChatOpen" id="open-chat">
                <img src="../assets/message_question_icon.svg" alt="chatbot">
            </span>
            <span v-else id="close-chat">&times;</span>
        </button>
        <div v-if="isChatOpen" class="chat-frame">
            <div class="chat-header">
                <h3>Chatbot</h3>
            </div>
            <div class="chat-messages" ref="chatContainer">
                <div v-for="(message, index) in messages" :key="index" :class="['message', message.sender]">
                    <span v-html="message.text"></span>
                </div>
            </div>
            <div v-if="isProcessing" class="processing-msg">
                <div class="globe"></div>
            </div>
            <div class="chat-input">
                <input v-model="userInput" @keyup.enter="sendMessage" placeholder="¿Qué horarios tiene el doctor?"
                    :disabled="isProcessing" />
                <button @click="sendMessage" :disabled="isProcessing">Enviar</button>
            </div>
        </div>
    </div>
</template>


<style>
.chatbot-container {
    position: fixed;
    bottom: 20px;
    right: 20px;
    z-index: 1000;
}

.chatbot-button {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 70px;
    height: 70px;
    border-radius: 50%;
    background-color: #3790D0;
    color: white;
    border: none;
    cursor: pointer;
    font-size: 16px;
    transition: 0.3s ease;
    will-change: transform;

    &:hover {
        transform: scale(1.1);
    }
}

.chatbot-button.open {
    background-color: #da3f34;
}

#open-chat {
    img {
        width: 35px;
    }
}

#close-chat {
    font-size: 54px;
}

.chat-frame {
    position: absolute;
    bottom: 80px;
    right: 0;
    width: 300px;
    height: 400px;
    border: 1px solid #ccc;
    border-radius: 10px;
    background-color: white;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
}

.chat-header {
    background-color: #d9dee9;
    color: #121215;
    padding: 10px;
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;
}

.chat-header h3 {
    margin: 0;
}

.chat-messages {
    flex-grow: 1;
    overflow-y: auto;
    padding: 10px;
}

.message {
    margin-bottom: 10px;
    padding: 8px 12px;
    border-radius: 1em;
    max-width: 80%;
    text-align: start;
}

.message.user {
    background-color: #e6f2ff;
    align-self: flex-end;
    margin-left: auto;
}

.message.bot {
    background-color: #f0f0f0;
    align-self: flex-start;
}

.chat-input {
    display: flex;
    padding: 10px;
    border-top: 1px solid #ccc;
}

.chat-input input {
    flex-grow: 1;
    padding: 5px;
    border: 1px solid #ccc;
    border-radius: 3px;
    font-family: poppins-regular;
}

.chat-input button {
    margin-left: 5px;
    padding: 5px 10px;
    background-color: #3790D0;
    font-family: poppins-regular;
    font-weight: 600;
    color: white;
    border: none;
    border-radius: 3px;
    cursor: pointer;
    transition: 0.25s;
}

.chat-input button:hover {
    background-color: #2176b3;
}

.processing-msg {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 40px;
}

.globe {
    width: 20px;
    height: 20px;
    border-radius: 50%;
    border: 2px solid #3790D0;
    border-top: 2px solid transparent;
    animation: spin 1s linear infinite;
}

@keyframes spin {
    0% {
        transform: rotate(0deg);
    }

    100% {
        transform: rotate(360deg);
    }
}
</style>