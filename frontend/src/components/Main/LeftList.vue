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
                   @click.stop="changeMainNav('dashboard')"
                   :active="mainStore.mainNav === 'dashboard'"></v-list-item>

      <v-list-item prepend-icon="mdi-application-cog-outline" value="messages"
                   @click.stop="changeMainNav('models')"
                   :active="mainStore.mainNav === 'models'"></v-list-item>

      <v-list-item prepend-icon="mdi-resistor-nodes" value="messages"
                   @click.stop="changeMainNav('custom_nodes')"
                   :active="mainStore.mainNav === 'custom_nodes'"></v-list-item>
    </v-list>
  </v-navigation-drawer>

</template>

<style scoped>

</style>