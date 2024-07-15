<script setup>
import {onMounted, onUnmounted, ref} from "vue";
import {Prompt, SaveFile} from "@wailsjs/go/backend/App.js";
import {useMainStore} from "@/stores/store.js";

const appButton = defineModel('appButton')
defineProps(['workspaceIndex'])

const snackbar = ref(false)
const timeout = 3000

const handlePrompt = async (event) => {
  if (event.data.type === 'saveFile') {
    const savePath = await Prompt('workflow.json', 'Save workflow as');

    const result = await SaveFile(savePath, event.data.json);

    if (result == null) {
      snackbar.value = true
    } else {
      // TODO
      alert("save fail")
      console.log("save fail", result)
    }
  }
  if (event.data.type === 'saveImg') {
    console.log("saveImg...", event.data)
  }
};

onMounted(() => {
  window.addEventListener('message', handlePrompt);
});
onUnmounted(() => {
  window.removeEventListener('message', handlePrompt);
});
</script>

<template>
  <div v-show="appButton === workspaceIndex">

    <template v-if="useMainStore().serverIsRunning">
      <iframe src="http://127.0.0.1:8189" class="comfy-screen"></iframe>
    </template>

    <template v-else>
      Server is not running, you should start server first.
    </template>

    <v-snackbar
        location="top"
        color="primary"
        v-model="snackbar"
        :timeout="timeout"
    >
      Save Workflow Success
      <template v-slot:actions>
        <v-btn
            color="green"
            variant="text"
            @click="snackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>

  </div>
</template>

<style scoped>
.comfy-screen {
  border: 0;
  display: block;
  width: 100%;
  height: calc(100vh - 40px);
}
</style>