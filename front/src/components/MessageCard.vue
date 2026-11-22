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
  <div>
    <h3>{{ message.callerName }} - <span>{{ message.callerId }}</span></h3>
    <audio ref="el" controls preload="none" :src="message.url" @play="emit('play', message.id)",
      @ended="emit('stop', message.id)"
    />
    <p>{{ message.created }}</p>
    <button>さくじょ</button>
  </div>
</template>
