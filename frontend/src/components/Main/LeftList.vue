<script setup>

import {useMainStore} from "@/stores/store.js";

const mainStore = useMainStore()

const changeMainNav = (nav) => {
  mainStore.setMainNav(nav)
}
</script>

<template>

  <v-navigation-drawer
      theme="dark"
      permanent
      rail
  >
    <v-list
        density="compact"
        nav
    >
      <v-list-item prepend-icon="mdi-view-dashboard" value="dashboard"
                   @click.stop="changeMainNav('dashboard'); mainStore.setCurrentPath('', '')"
                   :active="mainStore.mainNav === 'dashboard'"></v-list-item>

      <v-list-item prepend-icon="mdi-application-cog-outline" value="models"
                   @click.stop="changeMainNav('models')"
                   :active="mainStore.mainNav === 'models'"></v-list-item>

      <v-list-item prepend-icon="mdi-image-multiple-outline" value="output"
                   @click.stop="changeMainNav('output'); mainStore.setCurrentPath('output', '')"
                   :active="mainStore.mainNav === 'output'"></v-list-item>

      <v-list-item prepend-icon="mdi-resistor-nodes" value="custom_nodes"
                   @click.stop="changeMainNav('custom_nodes'); mainStore.setCurrentPath('custom_nodes', '')"
                   :active="mainStore.mainNav === 'custom_nodes'"></v-list-item>
    </v-list>
  </v-navigation-drawer>

</template>

<style scoped>

</style>