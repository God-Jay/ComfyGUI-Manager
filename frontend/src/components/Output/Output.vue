<script setup>
import {onActivated, ref} from "vue";
import {GetImages, GetImageWorkflow} from "@wailsjs/go/output/Output.js";
import {OpenFileInDir} from "@wailsjs/go/backend/App.js";
import moment from "moment";
import {EventsEmit} from "@wailsjs/runtime/runtime.js";
import {useMainStore} from "@/stores/store.js";

const images = ref([])

const getImages = () => {
  GetImages().then(r => {
    images.value = r
  })
}
getImages()

const workflowDetailDialog = ref(false)
const loadConfirmDialog = ref(false)
const clickFile = ref()
const clickIndex = ref()
const hasWorkflow = ref(false)
const workflow = ref('')
const viewWorkflow = (outputFile, i) => {
  GetImageWorkflow(outputFile.name).then(r => {
    if (r === '') {
      r = 'No workflow found'
      hasWorkflow.value = false
    } else {
      hasWorkflow.value = true
    }
    workflow.value = r
    workflowDetailDialog.value = true
    clickFile.value = outputFile
    clickIndex.value = i
  })
}

const loadWorkflow = (clickFile) => {
  EventsEmit("LoadImgWorkflow", workflow.value, clickFile.name)
  loadConfirmDialog.value = false
}

const offsetTop = ref(0)
const scrollTarget = ref(null)

const onScroll = (e) => {
  offsetTop.value = e.target.scrollTop
}

onActivated(() => {
  if (scrollTarget.value) {
    const element = scrollTarget.value.$el || scrollTarget.value.$refs.default; // 根据需要选择正确的引用
    if (element) {
      element.scrollTop = offsetTop.value;
    } else {
      console.error("Unable to access the DOM element");
    }
  }
});

const imgRefs = ref([])
const setImgRef = (i) => ($el) => {
  imgRefs.value[i] = $el
}

</script>

<template>

  <v-card
      elevation="0"
      class="mx-auto"
      min-width="860"
  >
    <v-container fluid
                 ref="scrollTarget"
                 id="container-scroll-target"
                 class="overflow-y-auto screen-height">
      <v-row dense v-if="images.length !== 0"
             v-scroll:#container-scroll-target="onScroll">
        <v-col
            v-for="(imageFile, i) in images"
            :key="i"
            cols="12" sm="3" md="2" lg="2" xl="1"
        >
          <v-card
              hover
              rounded="lg"
              class="mx-auto"
              max-width="344"
          >
            <!-- TODO port -->
            <v-img
                @click="OpenFileInDir('output', imageFile.name)"
                :src="'http://localhost:8190/output/'+imageFile.name"
                :ref="setImgRef(i)"
                cover
            ></v-img>
            <v-card-text class="text-center">
              <v-btn color="success" size="small" class="mb-2" @click.stop="viewWorkflow(imageFile, i)">
                view workflow
              </v-btn>
              <p>{{ imageFile.name }}</p>
              <p>{{ imgRefs[i].naturalHeight + '×' + imgRefs[i].naturalWidth + ' pix' }}</p>
              <p>{{ imageFile.size }}</p>
              <p>{{ moment(imageFile.modTime).format('LLL') }}</p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
      <!--TODO-->
      <template v-else>
        no output
      </template>
    </v-container>
  </v-card>

  <v-dialog
      attach="true"
      scroll-strategy="none"
      v-model="workflowDetailDialog"
      max-width="85%"
  >
    <v-card class="pa-1">
      <div class="d-flex flex-no-wrap justify-space-between">
        <div>
          <v-card-title>
            {{ clickFile.name }}
          </v-card-title>
          <v-card-text>
            {{ workflow }}
          </v-card-text>
        </div>

        <div class="text-center">
          <v-avatar
              class="ma-3"
              rounded="0"
              size="125"
          >
            <v-img :src="'http://localhost:8190/output/'+clickFile.name"></v-img>
          </v-avatar>
          <p>{{ imgRefs[clickIndex].naturalHeight + '×' + imgRefs[clickIndex].naturalWidth + ' pix' }}</p>
          <p>{{ clickFile.size }}</p>
          <p>{{ moment(clickFile.modTime).format('LLL') }}</p>
          <v-btn color="primary" @click="loadConfirmDialog = true" :disabled="!hasWorkflow">load workflow</v-btn>
        </div>
      </div>
    </v-card>
  </v-dialog>

  <v-dialog
      attach="true"
      scroll-strategy="none"
      v-model="loadConfirmDialog"
      width="auto"
  >
    <v-card
        max-width="400"
        prepend-icon="mdi-alert"
        text="Your current workflow will be overwritten. Are you sure you want to load?"
        title="load workflow"
    >
      <v-card-actions>
        <v-btn
            v-if="useMainStore().serverIsRunning"
            color="red"
            variant="flat"
            @click="loadWorkflow(clickFile)"
        >load
        </v-btn>
        <v-btn
            v-else
            color="red"
            variant="flat"
            disabled
        >server is not running
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

</template>

<style scoped>
.screen-height {
  height: calc(100vh - 80px);
}
</style>