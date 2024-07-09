<script setup>
import LeftList from "@/components/Main/LeftList.vue";
import Model from "@/components/Model/Model.vue";
import Dashboard from "@/components/Dashboard/Dashboard.vue";
import CustomNodes from "@/components/CustomNodes/CustomNodes.vue";
import {useMainStore} from "@/stores/store.js";

const mainStore = useMainStore()

</script>

<template>

  <LeftList/>

  <v-footer app>ComfyUI Path: {{ mainStore.comfyUIPath }}</v-footer>

  <Dashboard v-if="mainStore.mainNav === 'dashboard'"/>

  <Model v-if="mainStore.mainNav === 'models'"/>

  <CustomNodes v-if="mainStore.mainNav === 'custom_nodes'"/>

</template>

<style scoped>

</style>