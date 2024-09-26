<script setup>
import {onActivated} from "vue";
import Dir from "@/components/Model/LeftList/Dir.vue";
import {useModelStore} from "@/stores/model.js";
import {useMainStore} from "@/stores/store.js";

const mainStore = useMainStore()
const modelStore = useModelStore()


onActivated(() => {
  mainStore.setCurrentPath('models' + '/' + modelStore.selectedModelPath, '')
})
</script>

<template>

  <v-navigation-drawer permanent>
    <v-list-item title="Models" subtitle="model directory"
                 :class="{'bg-green': modelStore.isMainPath()}"
                 prepend-icon="mdi-file-multiple"
                 link
                 @click="modelStore.selectModelPath(''); mainStore.setCurrentPath('models' + '/' + modelStore.selectedModelPath, '')"
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