<script setup>
import LeftList from "@/components/Main/LeftList.vue";
import Model from "@/components/Model/Model.vue";
import Dashboard from "@/components/Dashboard/Dashboard.vue";
import CustomNodes from "@/components/CustomNodes/CustomNodes.vue";
import {ListModelDir} from "@wailsjs/go/models/Service.js";
import {useMainStore} from "@/stores/store.js";
import {useModelStore} from "@/stores/model.js";

const mainStore = useMainStore()
const modelStore = useModelStore()

async function getModels() {
  const modelDir = await ListModelDir()
  modelStore.setModelDir(modelDir)
}

getModels()

</script>

<template>

  <LeftList/>

  <v-footer app>ComfyUI Path: {{ mainStore.comfyUIPath }}</v-footer>

  <Dashboard v-if="mainStore.mainNav === 'dashboard'"/>

  <Model v-if="mainStore.mainNav === 'models'"/>

  <CustomNodes v-if="mainStore.mainNav === 'custom_nodes'"/>

</template>

<style scoped>

</style>