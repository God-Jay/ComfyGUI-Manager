<script setup>
import Fps from "@/Fps.vue";

const activeButtonClass = 'grey-lighten-3'
const inactiveButtonClass = 'grey'

const appButton = defineModel('appButton')
const workspaceNum = defineModel('workspaceNum')
</script>

<template>

  <v-app-bar :elevation="4" height="40" color="grey-darken-3">
    <Fps/>

    <v-btn variant="elevated" size="small" :color="`${appButton === -1 ? activeButtonClass : inactiveButtonClass}`"
           @click="appButton = -1"
           prepend-icon="mdi-home" class="ml-2">Main
    </v-btn>

    <template v-for="i in workspaceNum">
      <v-btn variant="elevated" size="small" :color="`${appButton === i ? activeButtonClass : inactiveButtonClass}`"
             @click="appButton = i"
             class="ml-2">Workspace #{{ i }}
      </v-btn>
    </template>

    <v-btn v-if="false"
           variant="elevated" size="small" class="ml-2" :color="inactiveButtonClass" @click="workspaceNum++">
      <v-icon icon="mdi-plus"></v-icon>
    </v-btn>
  </v-app-bar>

</template>

<style scoped>

</style>