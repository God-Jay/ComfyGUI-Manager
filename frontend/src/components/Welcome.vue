<script setup>
import {SelectFolder} from "@wailsjs/go/main/App.js";
import {ListModelDir} from '@wailsjs/go/models/Service'
import {useModelStore} from "@/stores/model.js";
import {useMainStore} from "@/stores/store.js";


const mainStore = useMainStore()
const modelStore = useModelStore()

async function selectFolder() {
  try {
    const folder = await SelectFolder('Select ComfyUI folder');
    console.log("selectFolder", folder)
    mainStore.setComfyUIPath(folder);
  } catch (error) {
    console.error('SelectFolder error', error);
  }
}

async function start() {
  const modelDir = await ListModelDir(mainStore.comfyUIPath)

  mainStore.setWelcome(false)
  modelStore.setModelDir(modelDir)
}

</script>

<template>

  <v-container class="fill-height">
    <v-row justify="center" align="center">
      <v-col cols="auto">
        <v-row justify="center">
          <v-text-field
              width="400"
              v-model="mainStore.comfyUIPath"
              label="ComfyUI folder"
              variant="solo-filled"
              disabled
          ></v-text-field>
        </v-row>
        <v-row justify="center">
          <v-btn @click="selectFolder" color="success">Open ComfyUI folder</v-btn>
        </v-row>
        <v-row justify="center" class="mt-6">
          <v-btn @click="start" color="primary">Start</v-btn>
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
