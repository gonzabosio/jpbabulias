<script setup>
import { ref, onMounted } from 'vue'

const isChatOpen = ref(false)
const messages = ref([])
const userInput = ref('')

const toggleChat = () => {
    isChatOpen.value = !isChatOpen.value
}

const sendMessage = () => {
    if (userInput.value.trim() === '') return

    messages.value.push({ text: userInput.value, sender: 'user' })
    userInput.value = ''

    // replace with actual chatbot logic
    setTimeout(() => {
        messages.value.push({ text: 'Ok', sender: 'bot' })
    }, 1000)
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
            <div class="chat-messages">
                <div v-for="(message, index) in messages" :key="index" :class="['message', message.sender]">
                    {{ message.text }}
                </div>
            </div>
            <div class="chat-input">
                <input v-model="userInput" @keyup.enter="sendMessage" placeholder="¿Qué horarios tiene el doctor?" />
                <button @click="sendMessage">Send</button>
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
    border-radius: 20px;
    max-width: 80%;
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
}

.chat-input button {
    margin-left: 5px;
    padding: 5px 10px;
    background-color: #3790D0;
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
</style>