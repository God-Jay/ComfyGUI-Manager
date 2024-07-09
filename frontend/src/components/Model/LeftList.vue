<script setup>

import folderImg from "@/assets/images/folder.png";
import {useModelStore} from "@/stores/model.js";

const modelStore = useModelStore()

</script>

<template>

  <v-navigation-drawer permanent>
    <v-list-item title="Models" subtitle="model directory"
                 :class="{'bg-green': 'main' === modelStore.modelNav}"
                 prepend-icon="mdi-file-multiple"
                 link @click="modelStore.setMainNav"
    ></v-list-item>
    <v-divider></v-divider>
    <template v-for="(dir, i) in modelStore.modelDir.subDirs">
      <v-list-item
          :class="{'bg-green': i === modelStore.modelIndex}"
          @click="modelStore.setDirIndex(dir.name, i)"
          link
          :title="`${dir.name} (${dir.fileCount})`"
          :prepend-avatar="folderImg"
      ></v-list-item>
    </template>
  </v-navigation-drawer>
</template>

<style scoped>

</style>