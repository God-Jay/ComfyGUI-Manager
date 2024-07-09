<script setup>
import {onBeforeUnmount, onMounted, ref} from "vue";

const fps = ref(0);
let lastFrameTime = performance.now();
let frameCount = 0;
let animationFrameId;

function updateFPS() {
  const currentFrameTime = performance.now();
  frameCount++;
  const deltaTime = currentFrameTime - lastFrameTime;

  if (deltaTime >= 1000) {
    fps.value = (frameCount / deltaTime) * 1000;
    frameCount = 0;
    lastFrameTime = currentFrameTime;
  }

  animationFrameId = requestAnimationFrame(updateFPS);
}

onMounted(() => {
  animationFrameId = requestAnimationFrame(updateFPS);
});

onBeforeUnmount(() => {
  cancelAnimationFrame(animationFrameId);
});
</script>

<template>
  <div>FPS: {{ fps.toFixed(2) }}</div>
</template>

<style scoped>

</style>