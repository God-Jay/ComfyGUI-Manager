<script setup>
import ComfyUIGitVersion from "@/components/Dashboard/Cards/ComfyUIGitVersion.vue";
import ComfyUIServerStatus from "@/components/Dashboard/Cards/ComfyUIServerStatus.vue";
import CustomNodes from "@/components/Dashboard/Cards/CustomNodes.vue";
</script>

<template>
  <v-container class="ma-0" style="max-width: 100%">
    <!--    comfy-->
    <v-row>
      <!--    comfy git-->
      <v-col cols="12" md="8">
        <ComfyUIGitVersion/>
      </v-col>

      <!--    comfy status-->
      <v-col cols="12" md="4">
        <ComfyUIServerStatus/>
      </v-col>

    </v-row>

    <v-row>
      <v-col cols="12" md="8">
        <CustomNodes/>
      </v-col>
    </v-row>

  </v-container>

</template>

<style scoped>
</style>