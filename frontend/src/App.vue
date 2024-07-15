<script setup>
import {ref} from "vue";
import Welcome from "@/components/Welcome.vue";
import Main from "@/components/Main/Main.vue";
import ComfyUI from "@/components/ComfyUI/ComfyUI.vue";
import AppBar from "@/components/AppBar.vue";
import {GetComfyUIPath} from "@wailsjs/go/backend/App.js";
import {useMainStore} from "@/stores/store.js";

const mainStore = useMainStore()

const appButton = ref(-1)

const workspaceNum = ref(1)

async function getComfyUIPath() {
  const path = await GetComfyUIPath();
  mainStore.setComfyUIPath(path);
}

getComfyUIPath()

</script>

<template>
  <v-app>
    <v-main>

      <template v-if="!mainStore.hasSetComfyUIPath">
        <Welcome/>
      </template>

      <template v-else>
        <AppBar v-model:appButton="appButton" v-model:workspaceNum="workspaceNum" style="z-index: 9999"/>

        <!--      Main-->
        <KeepAlive>
          <Main v-if="appButton === -1"/>
        </KeepAlive>

        <!--      Workspace-->
        <template v-for="i in workspaceNum">
          <KeepAlive>
            <ComfyUI v-model:appButton="appButton" :workspaceIndex="i"/>
          </KeepAlive>
        </template>
      </template>

    </v-main>
  </v-app>
</template>

<style>

</style>