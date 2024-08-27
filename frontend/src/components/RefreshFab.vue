<script setup>
import {ref} from "vue";

const props = defineProps({
  refresh: {
    type: Function,
    required: true,
  },
});


const spin = ref(false)

async function handleClick() {
  spin.value = true
  await props.refresh()
  spin.value = false
}

</script>

<template>
  <v-fab
      color="primary"
      :icon="spin ? 'mdi-refresh mdi-spin' : 'mdi-refresh'"
      class="mb-4"
      location="bottom end"
      absolute
      app
      @click.stop="handleClick"
  ></v-fab>
</template>

<style scoped>

</style>