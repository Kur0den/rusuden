<script setup lang="ts">
import { listMessages } from './api';
import { onMounted, ref } from 'vue';
import MessageCard from './components/MessageCard.vue';
import type { Message } from './type';

const playing = ref<number | null>(null)

const isLoading = ref(true)
const messages = ref<Message[]>([])
const errorMsg = ref("")

onMounted(async () => {
  try {
    messages.value = await listMessages()
  } catch(e) {
    errorMsg.value = e instanceof Error ? e.message : String(e)
  } finally {
    isLoading.value = false
  }
})
</script>

<template>
  <h1>rusuden</h1>
  <p v-if="isLoading">Loading...</p>
  <p v-else-if="errorMsg">{{ errorMsg }}</p>
  <p v-else-if="messages.length === 0">留守電はありません</p>
  <template v-else>
    <MessageCard v-for="m in messages" :key="m.id" :message="m" :playing="playing === m.id" @play="playing = m.id" @stop="playing = null" />
  </template>
</template>

<style scoped></style>
