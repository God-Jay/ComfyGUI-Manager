<script setup>
import {onBeforeUnmount, onMounted, ref} from "vue";
import {GetImages, GetImageWorkflow, GetStaredImages, StarImg} from "@wailsjs/go/output/Output.js";
import moment from "moment";
import {EventsEmit} from "@wailsjs/runtime/runtime.js";
import {useMainStore} from "@/stores/store.js";
import RefreshFab from "@/components/RefreshFab.vue";
import CardList from "@/components/Output/CardList.vue";

const tab = ref()

const images = ref([])
const staredImages = ref([])

const getImages = () => {
  images.value = []
  GetImages().then(r => {
    images.value = r
  })
  staredImages.value = []
  GetStaredImages().then(r => {
    staredImages.value = r
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

// TODO: improve performance
const setImgRef = (i, imageFile) => ($el) => {
  if (imageFile.ref !== undefined) {
    return
  }
  imageFile.ref = $el
}

const imgOverlay = ref(false)
const imgOverlayIndex = ref(0)

function clickImg(img, imgIndex) {
  imgOverlay.value = true
  imgOverlayIndex.value = imgIndex
}

const starImg = (img) => {
  StarImg(img.name).then(r => {
    img.stared = r
  })
}

const refresh = async () => {
  getImages()
};

const handleKeydown = (e) => {
  let imgLength;
  if (tab.value === "all-outputs") {
    imgLength = images.value.length
  } else {
    imgLength = staredImages.value.length
  }
  if (e.key === 'ArrowLeft') {
    if (imgOverlayIndex.value === 0) {
      imgOverlayIndex.value = imgLength - 1
    } else {
      imgOverlayIndex.value = imgOverlayIndex.value - 1
    }
  } else if (e.key === 'ArrowRight') {
    if (imgOverlayIndex.value === imgLength - 1) {
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

  <v-card class="screen-height">
    <v-tabs
        v-model="tab"
        align-tabs="center"
        color="deep-purple-accent-4"
        grow
    >
      <v-tab value="all-outputs">All Outputs</v-tab>
      <v-tab value="stared-outputs">Stars</v-tab>
    </v-tabs>

    <v-card-text class="pa-0">
      <v-tabs-window v-model="tab">
        <v-tabs-window-item value="all-outputs">
          <CardList :images="images" :clickImg="clickImg" :setImgRef="setImgRef" :starImg="starImg"
                    :viewWorkflow="viewWorkflow"/>
        </v-tabs-window-item>

        <v-tabs-window-item value="stared-outputs">
          <CardList :images="staredImages" :clickImg="clickImg" :setImgRef="setImgRef" :starImg="starImg"
                    :viewWorkflow="viewWorkflow"/>
        </v-tabs-window-item>

      </v-tabs-window>
    </v-card-text>
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
          <p>{{ clickFile.ref?.naturalWidth + '×' + clickFile.ref?.naturalHeight + ' pix' }}</p>
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
        <v-carousel-item v-for="(imageFile, i) in tab==='stared-outputs' ? staredImages : images" :value="i">
          <div class="d-flex fill-height justify-center align-center">
            <v-img
                @click.stop="imgOverlay = false"
                max-width="85wh"
                max-height="90vh"
                height="auto"
                :src="'http://localhost:8190/output/'+imageFile.name"
            >
              <v-toolbar color="transparent">
                <template v-slot:append>
                  <v-btn variant="tonal" icon="mdi-star" :color="imageFile.stared ? 'red': 'white'"
                         @click.stop="starImg(imageFile)"></v-btn>
                </template>
              </v-toolbar>
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