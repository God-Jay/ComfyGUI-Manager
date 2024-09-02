<script setup>
import folderImg from "@/assets/images/folder.png";
import {useModelStore} from "@/stores/model.js";
import {useMainStore} from "@/stores/store.js";

const mainStore = useMainStore()
const modelStore = useModelStore()

const dir = defineModel('dir')
const dirPath = defineModel('dirPath')
</script>

<template>
  <template v-if="dir.subDirs == null">
    <v-list-item
        :class="{'bg-green': dirPath === modelStore.selectedModelPath}"
        @click="modelStore.selectModelPath(dirPath); mainStore.setCurrentPath('models' + '/' + modelStore.selectedModelPath, '')"
        link
        :title="`${dir.name} (${dir.fileCount})`"
        :prepend-avatar="folderImg"
    ></v-list-item>
  </template>

  <template v-else>
    <v-list-group :value="dirPath">
      <template v-slot:activator="{ props }">
        <v-list-item
            :class="{'bg-green': dirPath === modelStore.selectedModelPath}"
            @click="modelStore.selectModelPath(dirPath); mainStore.setCurrentPath('models' + '/' + modelStore.selectedModelPath, '')"
            :prepend-avatar="folderImg"
            v-bind="props"
            :title="`${dir.name} (${dir.fileCount})`"
        ></v-list-item>
      </template>

      <template v-for="(subDir, i) in dir.subDirs">
        <Dir :dir="subDir" :dirPath="dirPath + '/' + subDir.name"></Dir>
      </template>
    </v-list-group>
  </template>
</template>

<style scoped>

</style>

<style lang="scss" scoped>
.v-list-group {
  --prepend-width: 8px;
}

</style>