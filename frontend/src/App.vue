<script setup lang="ts">
import { onMounted, ref } from 'vue'

const serverInput = ref('')
const numberInput = ref('')
const hash = ref('')
const password = ref('')
const isPasswordGenerated = ref(false)
const numberInputError = ref(false)
const match = ref(false)

function validateHandler() {
  const parsedNumber = parseInt(numberInput.value)
  numberInputError.value = isNaN(parsedNumber) || parsedNumber < 1 || parsedNumber >= 16

  if (!numberInputError.value) {
    generatePassword()
  } else {
    isPasswordGenerated.value = false
  }
}

async function generatePassword() {
  try {
    const response = await fetch('http://localhost:8080/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ length: parseInt(numberInput.value) })
    })

    if (!response.ok) {
      throw new Error('Error generating password')
    }

    const data = await response.json()

    serverInput.value = data.Password
    hash.value = data.Hash
    password.value = data.Password
    match.value = data.Match

    isPasswordGenerated.value = true
  } catch (error) {
    console.error(error)
  }
}

function setServerValue(value: string) {
  serverInput.value = value
}

function copyPassword() {
  navigator.clipboard.writeText(serverInput.value)
}

const simulateInput = () => {
  numberInput.value = ''
}

const handleNumberInput = (event: Event) => {
  const input = (event.target as HTMLInputElement).value
  numberInput.value = input.replace(/[^0-9]/g, '')
}

onMounted(simulateInput)
</script>

<template>
  <div
    class="min-h-screen text-white bg-gradient-to-r from-blue-950 to-slate-900 flex flex-col items-center justify-center"
  >
    <h1 class="text-4xl mb-4 mr-3 flex">
      <img src="../public/go-gopher.svg" class="size-16" alt="gopher" />
      <p class="mt-3">GoPassword</p>
    </h1>
    <div class="flex flex-row items-center">
      <input
        type="text"
        placeholder="Generated password"
        class="bg-inherit border border-white text-white text-sm rounded-lg block w-full p-2.5"
        :value="serverInput"
        @input="setServerValue($event.target ? ($event.target as HTMLInputElement).value : '')"
        readonly
      />
      <button
        class="mt-1.5 ml-2 text-white bg-blue-500 hover:bg-blue-700 font-medium px-5 py-2.5 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 text-sm rounded-lg block w-full p-2.5"
        @click="copyPassword"
      >
        Copy password
      </button>
    </div>
    <div class="flex flex-row items-center mt-2">
      <input
        placeholder="Character count"
        class="mb-1.5 bg-inherit border border-white text-sm text-white rounded-lg block w-full p-2.5"
        v-model="numberInput"
        @input="handleNumberInput"
      />
      <button
        class="ml-2 text-white bg-blue-500 hover:bg-blue-700 font-medium px-5 py-2.5 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 text-sm rounded-lg block w-full p-2.5"
        @click="validateHandler"
      >
        Enter
      </button>
    </div>

    <p v-if="numberInputError" class="text-red-500 text-sm mt-1">
      Character count must be a non-empty number less than 16.
    </p>
    <span v-if="isPasswordGenerated" class="text-white mt-2">
      <p>Hash: {{ hash }}</p>
      <p class="mt-1.5 text-center">
        Password match:
        <span :class="match ? 'bg-green-600' : 'bg-red-500'" class="px-2 py-0.5 rounded-lg">
          {{ match }}
        </span>
      </p>
    </span>
  </div>
</template>
