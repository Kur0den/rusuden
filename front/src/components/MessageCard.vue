<script setup lang="ts">
import type { Message } from "@/type"
import { ref, watch } from "vue";
const props = defineProps<{message: Message; playing: boolean}>()
const emit = defineEmits<{play: [id: number]; stop: [id: number]}>()

const el = ref<HTMLAudioElement>()

watch(() => props.playing, (on) => {
  if (!on) el.value?.pause()
})
</script>

<template>
  <div class="card">
    <h3 class="caller-name">{{ message.callerName }}</h3> - <h4 class="caller-id">{{ message.callerId }}</h4>
    <audio rel="el" controls preload="none" :src="message.url" @play="emit('play', message.id)",
      @ended="emit('stop', message.id)"
    />
  </div>
</template>
