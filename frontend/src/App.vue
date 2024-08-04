<script setup lang="ts">
import { ref } from 'vue'
const serverInput = ref('')
const numberInput = ref('')
const hash = ref('')
const password = ref('')
const isPasswordGenerated = ref(false)

function setServerInput(value: string) {
  serverInput.value = value
}

function copyServerInput() {
  navigator.clipboard.writeText(serverInput.value)
  isPasswordGenerated.value = true
}

function checkHash() {}

</script>

<template>
  <div class="min-h-screen text-white bg-gradient-to-r from-blue-950 to-slate-900 flex flex-col items-center justify-center">
    <h1 class="text-4xl mb-4 mr-3 flex"><img src='../public/go-gopher.svg' class="size-16 " alt="gopher" /><p class="mt-3">GoPassword</p></h1>
    <div class="flex flex-row items-center">
      <input
        type="text"
        placeholder="Generated password"
        class="bg-inherit border border-white text-white text-sm rounded-lg block w-full p-2.5"
        :value="serverInput"
        @input="setServerInput($event.target ? ($event.target as HTMLInputElement).value : '')"
        readonly
      />
      <button
        class="mt-1.5 ml-2 text-white bg-blue-500 hover:bg-blue-700 font-medium px-5 py-2.5  mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 text-sm rounded-lg block w-full p-2.5"
        @click="copyServerInput"
      >
        Copy password
      </button>
    </div>
    <div class="flex flex-row items-center mt-2">
      <input
        placeholder="Character count"
        class="mb-1.5 bg-inherit border border-white text-sm text-white rounded-lg block w-full p-2.5"
        v-model="numberInput"
      />
      <button
        class="ml-2 text-white bg-blue-500 hover:bg-blue-700 font-medium px-5 py-2.5  mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 text-sm rounded-lg block w-full p-2.5"
        @click="copyServerInput"
      >
        Enter
      </button>
    </div>
    <span v-if="isPasswordGenerated" class="text-white mt-2">
    <p>Hash: {{ hash }}</p>
    <p>Password match: <span :class="hash === password ? 'bg-green-600' : 'bg-red-500'" class="px-2 py-0.5 rounded-lg">{{ hash === password }}</span></p>
  </span>
  </div>
</template>