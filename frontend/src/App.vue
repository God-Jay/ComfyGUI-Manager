<script setup>
import {ref} from "vue";
import Welcome from "@/components/Welcome.vue";
import Main from "@/components/Main/Main.vue";
import ComfyUI from "@/components/ComfyUI/ComfyUI.vue";
import AppBar from "@/components/AppBar.vue";
import {useMainStore} from "@/stores/store.js";

const mainStore = useMainStore()

const appButton = ref(-1)

const workspaceNum = ref(1)

</script>

<template>
  <v-app>
    <v-main>

      <template v-if="mainStore.welcome">
        <Welcome/>
      </template>

      <template v-else>
        <AppBar v-model:appButton="appButton" v-model:workspaceNum="workspaceNum"/>

        <!--      Main-->
        <template v-if="appButton === -1">
          <Main/>
        </template>

        <!--      Workspace-->
        <template v-for="i in workspaceNum">
          <ComfyUI v-if="appButton === i"/>
        </template>
      </template>

    </v-main>
  </v-app>
</template>

<style>

</style>