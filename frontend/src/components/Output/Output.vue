<script setup>
import {onActivated, onBeforeUnmount, onMounted, ref} from "vue";
import {GetImages, GetImageWorkflow} from "@wailsjs/go/output/Output.js";
import {OpenFileInDir} from "@wailsjs/go/backend/App.js";
import moment from "moment";
import {EventsEmit} from "@wailsjs/runtime/runtime.js";
import {useMainStore} from "@/stores/store.js";
import RefreshFab from "@/components/RefreshFab.vue";

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
    const element = scrollTarget.value.$el || scrollTarget.value.$refs.default;
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

const imgOverlay = ref(false)
const imgOverlayIndex = ref(0)

function clickImg(img, imgIndex) {
  imgOverlay.value = true
  imgOverlayIndex.value = imgIndex
}

const refresh = async () => {
  offsetTop.value = 0
  images.value = []
  getImages()
};

const handleKeydown = (e) => {
  if (e.key === 'ArrowLeft') {
    if (imgOverlayIndex.value === 0) {
      imgOverlayIndex.value = imgRefs.value.length - 1
    } else {
      imgOverlayIndex.value = imgOverlayIndex.value - 1
    }
  } else if (e.key === 'ArrowRight') {
    if (imgOverlayIndex.value === imgRefs.value.length - 1) {
      imgOverlayIndex.value = 0
    } else {
      imgOverlayIndex.value = imgOverlayIndex.value + 1
    }
  }
};

onMounted(() => {
  window.addEventListener('keydown', handleKeydown);
});

onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleKeydown);
});

</script>

<template>

  <RefreshFab :refresh="refresh"></RefreshFab>

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
                aspect-ratio="1"
                @click.stop="clickImg(imageFile, i)"
                :src="'http://localhost:8190/output/'+imageFile.name"
                :ref="setImgRef(i)"
            ></v-img>
            <v-card-text class="text-center pt-1 pb-1">
              <v-btn color="primary" size="small" class="mb-2" @click="OpenFileInDir('output', imageFile.name)">
                open in path
              </v-btn>
              <v-btn color="success" size="small" class="mb-2" @click.stop="viewWorkflow(imageFile, i)">
                view workflow
              </v-btn>
              <p>{{ imageFile.name }}</p>
              <p>{{ imgRefs[i].naturalWidth + '×' + imgRefs[i].naturalHeight + ' pix' }}</p>
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

  <!--  view workflow-->
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
              size="180"
          >
            <v-img :src="'http://localhost:8190/output/'+clickFile.name"></v-img>
          </v-avatar>
          <p>{{ imgRefs[clickIndex].naturalWidth + '×' + imgRefs[clickIndex].naturalHeight + ' pix' }}</p>
          <p>{{ clickFile.size }}</p>
          <p>{{ moment(clickFile.modTime).format('LLL') }}</p>
          <v-btn color="primary" @click="loadConfirmDialog = true" :disabled="!hasWorkflow">load workflow</v-btn>
        </div>
      </div>
    </v-card>
  </v-dialog>

  <!--  load workflow-->
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

  <!--  img popup-->
  <v-overlay v-model="imgOverlay"
             attach="true"
             class="align-center justify-center"
             width="95%">
    <v-container style="max-width: 100%" class="screen-height">
      <v-carousel hide-delimiters progress="primary" height="100%" v-model="imgOverlayIndex">
        <v-carousel-item v-for="(imageFile, i) in images" :value="i">
          <div class="d-flex fill-height justify-center align-center">
            <v-img
                @click.stop="imgOverlay = false"
                max-width="85wh"
                max-height="90vh"
                height="auto"
                :src="'http://localhost:8190/output/'+imageFile.name"
            >
              <template v-slot:placeholder>
                <div class="d-flex align-center justify-center fill-height">
                  <v-progress-circular
                      color="grey-lighten-4"
                      indeterminate
                  ></v-progress-circular>
                </div>
              </template>
            </v-img>
          </div>
        </v-carousel-item>
      </v-carousel>
    </v-container>
  </v-overlay>

</template>

<style scoped>
.screen-height {
  height: calc(100vh - 80px);
}
</style>