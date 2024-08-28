<script setup>
import folderImg from "@/assets/images/folder.png";
import {useModelStore} from "@/stores/model.js";

const modelStore = useModelStore()

const dir = defineModel('dir')
const dirPath = defineModel('dirPath')
</script>

<template>
  <template v-if="dir.subDirs == null">
    <v-list-item
        :class="{'bg-green': dirPath === modelStore.selectedModelPath}"
        @click="modelStore.selectModelPath(dirPath)"
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
            @click="modelStore.selectModelPath(dirPath)"
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