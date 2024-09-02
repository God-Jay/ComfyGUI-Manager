<script setup>
import LeftList from "@/components/Main/LeftList.vue";
import Model from "@/components/Model/Model.vue";
import Output from "@/components/Output/Output.vue";
import Dashboard from "@/components/Dashboard/Dashboard.vue";
import CustomNodes from "@/components/CustomNodes/CustomNodes.vue";
import {ListModelDir} from "@wailsjs/go/models/Service.js";
import {OpenFileInDir, OpenFolder} from "@wailsjs/go/backend/App.js";
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

  <LeftList style="z-index: 9998"/>

  <v-footer app style="justify-content: space-between;">
    <v-hover v-slot="{ isHovering, props }">
      <p @click="OpenFolder"
         v-bind="props"
         style="cursor: pointer"
         :class="isHovering ? 'text-primary' : 'text-secondary'">
        ComfyUI Path: {{ mainStore.comfyUIPath }}
      </p>
    </v-hover>
    <v-hover v-slot="{ isHovering, props }">
      <p @click="OpenFileInDir(mainStore.currentPath.path, mainStore.currentPath.fileName)"
         v-bind="props"
         style="cursor: pointer"
         :class="isHovering ? 'text-primary' : 'text-secondary'">
        Open Current Path
      </p>
    </v-hover>
  </v-footer>

  <KeepAlive>
    <Dashboard v-if="mainStore.mainNav === 'dashboard'"/>
  </KeepAlive>

  <KeepAlive>
    <Model v-if="mainStore.mainNav === 'models'"/>
  </KeepAlive>

  <KeepAlive>
    <Output v-if="mainStore.mainNav === 'output'"/>
  </KeepAlive>

  <KeepAlive>
    <CustomNodes v-if="mainStore.mainNav === 'custom_nodes'"/>
  </KeepAlive>

</template>

<style scoped>

</style>