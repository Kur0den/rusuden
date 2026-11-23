<script setup lang="ts">
import { listMessages } from './api'
import { onMounted, ref } from 'vue'
import MessageCard from './components/MessageCard.vue'
import type { Message } from './type'

const playing = ref<number | null>(null)

const isLoading = ref(true)
const messages = ref<Message[]>([])
const errorMsg = ref('')

onMounted(async () => {
  try {
    messages.value = await listMessages()
  } catch (e) {
    errorMsg.value = e instanceof Error ? e.message : String(e)
  } finally {
    isLoading.value = false
  }
})

function deleteMessage(id: number) {
  console.log('Delete: ' + id)
}
</script>

<template class="min-h-screen bg-base-100 text-base-content">
  <div class="mx-auto max-w-2xl px-5 py-8">
    <h1 class="mb-6 text-2xl font-bold tracking-wide">rusuden</h1>
    <p v-if="isLoading" class="py-8 text-center opacity-60">Loading...</p>
    <p v-else-if="errorMsg" class="alert alert-error">{{ errorMsg }}</p>
    <p v-else-if="messages.length === 0" class="py-8 text-center opacity-60">留守電はありません</p>
    <template v-else class="flex flex-col gap-3">
      <MessageCard
        v-for="m in messages"
        :key="m.id"
        :message="m"
        :playing="playing === m.id"
        @play="playing = m.id"
        @stop="playing = null"
        @delete="deleteMessage"
      />
    </template>
  </div>
</template>

<style scoped></style>
