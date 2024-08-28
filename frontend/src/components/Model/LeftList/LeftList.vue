<script setup>
import Dir from "@/components/Model/LeftList/Dir.vue";
import {useModelStore} from "@/stores/model.js";

const modelStore = useModelStore()

</script>

<template>

  <v-navigation-drawer permanent>
    <v-list-item title="Models" subtitle="model directory"
                 :class="{'bg-green': modelStore.isMainPath()}"
                 prepend-icon="mdi-file-multiple"
                 link @click="modelStore.selectModelPath('')"
    ></v-list-item>
    <v-divider></v-divider>
    <v-list>
      <template v-for="(dir, i) in modelStore.modelDir.subDirs">

        <Dir :dir="dir" :dirPath="dir.name"></Dir>

      </template>
    </v-list>
  </v-navigation-drawer>
</template>

<style scoped>

</style>