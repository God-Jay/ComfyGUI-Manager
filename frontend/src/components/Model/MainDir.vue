<script setup>
import folderImg from "@/assets/images/folder.png";
import {useModelStore} from "@/stores/model.js";
import {ref} from "vue";
import RefreshFab from "@/components/RefreshFab.vue";
import {ListModelDir} from "@wailsjs/go/models/Service.js";

const modelStore = useModelStore()

const dialog = ref(false)

const dirInfo = ref()

const clickCard = (dir) => {
  if (dir === dirInfo.value) {
    dialog.value = !dialog.value
  } else {
    dialog.value = true
    dirInfo.value = dir
  }
}

const refresh = async () => {
  const modelDir = await ListModelDir()
  modelStore.setModelDir(modelDir)
};
</script>

<template>

  <RefreshFab :refresh="refresh"></RefreshFab>

  <v-card
      elevation="0"
      class="mx-auto"
      min-width="600"
  >
    <v-container fluid
                 class="overflow-y-auto screen-height">
      <v-row dense>
        <v-col
            v-for="(dir, i) in modelStore.modelDir.subDirs"
            :key="i"
            cols="12" sm="3" md="2" lg="2" xl="1"
        >
          <v-tooltip :text="`File count: ${dir.fileCount}`">
            <template v-slot:activator="{ props }">
              <v-card
                  @click.stop="clickCard(dir)"
                  v-bind="props"
                  hover
                  rounded="lg"
                  class="mx-auto mb-1"
                  max-width="344"
              >
                <v-img
                    :src="folderImg"
                    cover
                ></v-img>
                <v-card-text class="text-center pt-0 pb-2">
                  <p>{{ dir.name }}</p>
                  <p>{{ dir.totalFileSize }}</p>
                </v-card-text>
              </v-card>
            </template>
          </v-tooltip>

        </v-col>
      </v-row>
    </v-container>
  </v-card>

  <v-dialog
      attach="true"
      scroll-strategy="none"
      v-model="dialog"
      width="auto"
  >
    <v-card
        max-width="400"
        prepend-icon="mdi-update"
        :title="dirInfo.name"
    >
      <v-card-text>
        <div>File count in this directory: {{ dirInfo.fileCount }}</div>
      </v-card-text>
      <v-card-text>
        <div>Total file size in this directory: {{ dirInfo.totalFileSize }}</div>
      </v-card-text>
      <template v-slot:actions>
        <v-btn
            class="ms-auto"
            text="Ok"
            @click="dialog = false"
        ></v-btn>
      </template>
    </v-card>
  </v-dialog>

</template>

<style scoped>
.screen-height {
  height: calc(100vh - 80px);
}
</style>