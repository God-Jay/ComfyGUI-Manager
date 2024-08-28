<script setup>
import MainDir from "@/components/Model/MainDir.vue";
import ModelDir from "@/components/Model/ModelDir.vue";
import LeftList from "@/components/Model/LeftList/LeftList.vue";
import {useModelStore} from "@/stores/model.js";

const modelStore = useModelStore()
</script>

<template>

  <LeftList/>

  <MainDir v-if="modelStore.isMainPath()"/>

  <ModelDir v-if="!modelStore.isMainPath()"/>

</template>

<style scoped>

</style>

