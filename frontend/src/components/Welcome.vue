<script setup>
import {ref} from "vue";
import {SelectFolder, SetComfyUIPath} from "@wailsjs/go/backend/App.js";
import {useMainStore} from "@/stores/store.js";

const mainStore = useMainStore()

const selectedComfyUIPath = ref('')

async function selectFolder() {
  try {
    selectedComfyUIPath.value = await SelectFolder('Select ComfyUI folder');
    console.log("selectFolder", selectedComfyUIPath)
  } catch (error) {
    console.error('SelectFolder error', error)
  }
}

async function start() {
  await SetComfyUIPath(selectedComfyUIPath.value)
  mainStore.setComfyUIPath(selectedComfyUIPath.value)
}

</script>

<template>

  <v-container class="fill-height">
    <v-row justify="center" align="center">
      <v-col cols="auto">
        <v-row justify="center">
          <v-text-field
              width="400"
              v-model="selectedComfyUIPath"
              label="ComfyUI folder"
              variant="solo-filled"
              disabled
          ></v-text-field>
        </v-row>
        <v-row justify="center">
          <v-btn @click="selectFolder" color="success">Open ComfyUI folder</v-btn>
        </v-row>
        <v-row justify="center" class="mt-6">
          <v-btn @click="start" color="primary" :disabled="selectedComfyUIPath === ''">Start</v-btn>
        </v-row>
      </v-col>
    </v-row>
  </v-container>

</template>

<style lang="scss" scoped>
.v-btn {
  text-transform: none;
}
</style>
