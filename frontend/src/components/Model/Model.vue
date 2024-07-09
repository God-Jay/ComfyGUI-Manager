<script setup>
import MainDir from "@/components/Model/MainDir.vue";
import ModelDir from "@/components/Model/ModelDir.vue";
import LeftList from "@/components/Model/LeftList.vue";
import {useModelStore} from "@/stores/model.js";

const modelStore = useModelStore()
</script>

<template>

  <LeftList/>

  <MainDir v-if="modelStore.modelNav === 'main'"/>

  <ModelDir v-if="modelStore.modelNav !== 'main'"/>

</template>

<style scoped>

</style>

