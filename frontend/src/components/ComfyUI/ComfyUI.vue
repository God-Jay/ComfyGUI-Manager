<script setup>
import {onMounted, onUnmounted, ref} from "vue";
import {Prompt, SaveFile} from "@wailsjs/go/backend/App.js";
import {EventsOn} from "@wailsjs/runtime/runtime.js";
import {StoreOutput} from "@wailsjs/go/comfyUI/ComfyUI.js";
import {useMainStore} from "@/stores/store.js";

const iframeRef = ref()
const iframeSrc = ref(`http://127.0.0.1:8189?cacheBuster=${new Date().getTime()}`)

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

  if (event.data.type === 'IFRAME_SAVE_OUTPUT') {
    StoreOutput(event.data.value, outputFileName.value)
  }
};

const outputFileName = ref('')
EventsOn("OutputImageFile", (output) => {
  outputFileName.value = output
  iframeRef.value.contentWindow.postMessage({type: 'VUE_SAVE_OUTPUT'}, "*");
})
EventsOn("LoadImgWorkflow", (workflow, name) => {
  iframeRef.value.contentWindow.postMessage({type: 'LOAD_IMG_WORKFLOW', workflow: workflow, name: name}, "*");
})

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
      <iframe ref="iframeRef" :src="iframeSrc" class="comfy-screen"></iframe>
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