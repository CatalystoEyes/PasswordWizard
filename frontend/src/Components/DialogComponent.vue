<script setup>
import { ref } from 'vue'
import { TransitionRoot, TransitionChild, Dialog, DialogPanel, DialogTitle } from '@headlessui/vue'
import { Heart } from 'lucide-vue-next'
import { X } from 'lucide-vue-next'

const isOpen = ref(false)

function closeModal() {
  isOpen.value = false
}
function openModal() {
  isOpen.value = true
}
</script>

<template>
  <div class="">
    <button
      type="button"
      @click="openModal"
      class="rounded-md bg-blue-600 ml-2 px-6 py-2.5 mb-2 text-sm font-medium text-white hover:bg-blue-700 focus:outline-none focus-visible:ring-2 focus-visible:ring-white/75"
    >
      Enter
    </button>
  </div>
  <TransitionRoot appear :show="isOpen" as="template">
    <Dialog as="div" @close="closeModal" class="relative z-10">
      <TransitionChild
        as="template"
        enter="duration-300 ease-out"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="duration-200 ease-in"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-black/25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4 text-center">
          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
            <DialogPanel
              class="w-full max-w-md transform overflow-hidden rounded-2xl bg-slate-200 p-6 text-left align-middle shadow-xl transition-all"
            >
              <DialogTitle as="h3" class="text-lg font-medium leading-6 text-gray-900">
                Generated password
              </DialogTitle>
              <div class="mt-2">
                <div class="flex flex-row items-center">
                  <p>{{ text }}</p>
                </div>
              </div>

              <div class="mt-4 flex">
                <button
                  type="button"
                  class="ml-2 text-white bg-blue-500 hover:bg-blue-700 font-medium px-5 py-2.5 mb-2 text-sm rounded-lg w-full select-none flex items-center"
                  @click="closeModal"
                >
                  <Heart class="mr-2" />Add to favourite
                </button>
                <button
                  type="button"
                  class="ml-2 text-white bg-blue-500 hover:bg-blue-700 font-medium px-5 py-2.5 mb-2 text-sm rounded-lg w-full select-none flex items-center"
                  @click="closeModal"
                >
                  <X class="mr-2" />Close tab
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
