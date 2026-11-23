<script setup lang="ts">
import type { Message } from '@/type'
import { ref, watch } from 'vue'
const props = defineProps<{ message: Message; playing: boolean }>()
const emit = defineEmits<{ play: [id: number]; stop: [id: number]; delete: [id: number] }>()

const el = ref<HTMLAudioElement>()

watch(
  () => props.playing,
  (on) => {
    if (!on) el.value?.pause()
  },
)
</script>

<template>
  <div class="card border border-base-300 bg-base-200 shadow-sm">
    <div class="card-body gap-2 p-4">
      <h3 class="flex items-baseline justify-self-start gap-3 text-xl font-bold">
        {{ message.callerName }}  <span class="text-sm font-semibold">{{ message.callerId }}</span>
      </h3>
      <audio
        ref="el"
        controls
        preload="none"
        :src="message.url"
        @play="emit('play', message.id)"
        @ended="emit('stop', message.id)"
      />
      <div class="flex item-center justify-between gap-3">
        <p class="text-sm opacity-60">{{ message.created }}</p>
        <button type="button" class="btn btn-outline btn-error btn-sm" @click="emit('delete', message.id)">さくじょ</button>
      </div>
    </div>
  </div>
</template>
