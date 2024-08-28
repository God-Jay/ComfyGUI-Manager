<script setup>
import {onActivated, ref} from "vue";
import modelImg from "@/assets/images/model.png";
import {GetModelFileSHA256} from "@wailsjs/go/models/Service";
import {BrowserOpenURL} from "@wailsjs/runtime/runtime.js";
import {getCivitaiInfo} from "@/api/civitai.js";
import {useModelStore} from "@/stores/model.js";
import moment from "moment";
import RefreshFab from "@/components/RefreshFab.vue";
import {ListModelDir} from "@wailsjs/go/models/Service.js";

const modelStore = useModelStore()

const state = ref()
const failure = ref(false)

const dialog = ref(false)
const fileInfo = ref()

const fileSHA256 = ref('')

async function clickModel(file) {
  if (file === fileInfo.value) {
    dialog.value = !dialog.value
  } else {
    failure.value = false
    state.value = undefined

    dialog.value = true
    fileInfo.value = file

    if (fileInfo.value.civitaiInfo !== undefined) {
      return
    }
    state.value = "getting model file sha256..."
    fileSHA256.value = await GetModelFileSHA256(file.path)

    try {
      state.value = "fetching model file info from civitai..."
      fileInfo.value.civitaiInfo = await getCivitaiInfo(fileSHA256.value)
    } catch (err) {
      failure.value = true
      if (err.response !== undefined && err.response.status === 404 && err.response.data?.error !== undefined) {
        state.value = "fetch civitai info error: " + err.response?.data?.error
      } else {
        state.value = "fetch civitai info error: " + err.message
      }
    }
  }
}

const imgOverlay = ref(false)
const imgOverlayIndex = ref(0)

function clickImg(img, imgIndex) {
  imgOverlay.value = true
  imgOverlayIndex.value = imgIndex
}

function replaceLargeImg(image) {
  const url = image.url
  const width = image.width
  return url.replace(/width=\d+/g, `width=${width}`)
}

function openCivitai(url) {
  BrowserOpenURL(url)
}

const currentImgModel = ref(0)
const snackbar = ref(false)
const copyText = async (text) => {
  try {
    await navigator.clipboard.writeText(text)
    snackbar.value = true
  } catch (err) {
    console.error('Failed to copy text:', err)
  }
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

const refresh = async () => {
  offsetTop.value = 0
  const modelDir = await ListModelDir()
  modelStore.setModelDir(modelDir)

  const dir = modelStore.modelDir.subDirs[modelStore.modelIndex]
  modelStore.setDirIndex(dir.name, modelStore.modelIndex)
};
</script>

<template>

  <RefreshFab :refresh="refresh"></RefreshFab>

  <!--  model list-->
  <v-card
      elevation="0"
      class="mx-auto"
      min-width="860"
  >
    <v-container fluid
                 ref="scrollTarget"
                 id="container-scroll-target"
                 class="overflow-y-auto screen-height">
      <v-row dense
             v-scroll:#container-scroll-target="onScroll">
        <v-col
            v-for="(file, i) in modelStore.specificModelDir.files"
            :key="i"
            cols="12" sm="3" md="2" lg="2" xl="1"
        >
          <v-card
              @click.stop="clickModel(file)"
              hover
              rounded="lg"
              class="mx-auto"
              max-width="344"
          >
            <v-img
                :src="modelImg"
                cover
            ></v-img>
            <v-card-text class="text-center">
              <p>{{ file.name }}</p>
              <p>{{ file.size }}</p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-card>

  <!--  model popup-->
  <v-dialog
      attach="true"
      scroll-strategy="none"
      v-model="dialog"
      width="auto"
  >
    <v-card min-width="700" :title="`${fileInfo.name}  [${fileInfo.size}]`">
      <div class="d-flex flex-no-wrap justify-space-between" v-if="fileInfo.civitaiInfo">
        <div>
          <v-card-title class="text-h6">
            Model Name:
            <v-chip>{{ fileInfo.civitaiInfo.model.name }}</v-chip>
          </v-card-title>
          <v-card-subtitle>
            BaseModel:
            <v-chip label color="primary" class="ml-1 mr-4">{{ fileInfo.civitaiInfo.baseModel }}</v-chip>
            Type:
            <v-chip label color="primary" class="ml-1">{{ fileInfo.civitaiInfo.model.type }}</v-chip>
            Published At:
            <v-chip label color="primary" class="ml-1">{{ moment(fileInfo.civitaiInfo.publishedAt).format('LLL') }}
            </v-chip>
          </v-card-subtitle>

          <v-card-text>
            Civitai:
            <v-btn
                variant="flat"
                color="success"
                size="small"
                @click="openCivitai(`https://civitai.com/models/${fileInfo.civitaiInfo.modelId}`)"
                target="_blank">View in Civitai
            </v-btn>
          </v-card-text>


          <v-card color="grey"
                  title="Prompt"
                  class="ma-2" max-width="900">
            <template v-slot:append>
              <v-btn variant="plain" icon="mdi-content-copy" size="sm"
                     @click="copyText(fileInfo.civitaiInfo.images[currentImgModel]?.meta?.prompt)"></v-btn>
            </template>
            <v-card-text>{{ fileInfo.civitaiInfo.images[currentImgModel]?.meta?.prompt }}</v-card-text>
          </v-card>

          <v-card color="grey"
                  title="Negative Prompt"
                  class="ma-2" max-width="900">
            <template v-slot:append>
              <v-btn variant="plain" icon="mdi-content-copy" size="sm"
                     @click="copyText(fileInfo.civitaiInfo.images[currentImgModel]?.meta?.negativePrompt)"></v-btn>
            </template>
            <v-card-text>{{ fileInfo.civitaiInfo.images[currentImgModel]?.meta?.negativePrompt }}</v-card-text>
          </v-card>

          <div v-html="fileInfo.civitaiInfo.description" class="pl-8" style="max-width: 600px"></div>
        </div>

        <!--        img carousel-->
        <div class="float-right mr-3" style="width: 400px">
          <v-carousel hide-delimiters progress="success" v-model="currentImgModel">
            <template v-for="(image, i) in fileInfo.civitaiInfo.images">
              <v-carousel-item>
                <div class="d-flex justify-space-around align-center bg-grey-lighten-4">
                  <v-img
                      @click.stop="clickImg(image, i)"
                      :aspect-ratio="1"
                      class="bg-white v-card--link"
                      :src="image.url"
                      width="400"
                      height="500"
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
            </template>
          </v-carousel>
        </div>
      </div>
      <div class="text-center" v-else>
        <v-progress-circular v-if="!failure" color="primary" indeterminate></v-progress-circular>
        <v-icon v-if="failure" color="red" size="32">mdi-alert</v-icon>
        <div class="text-subtitle-2 text-capitalize">{{ state }}</div>
      </div>
      <template v-slot:actions>
        <v-btn
            class="ms-auto"
            text="Ok"
            @click="dialog = false"
        ></v-btn>
      </template>
    </v-card>
  </v-dialog>

  <!--  img popup-->
  <v-overlay v-model="imgOverlay"
             attach="true"
             class="align-center justify-center"
             width="100%">
    <v-container style="max-width: 100%">
      <v-carousel hide-delimiters progress="success" height="100vh" v-model="imgOverlayIndex">
        <v-carousel-item v-for="(image, i) in fileInfo.civitaiInfo.images" :value="i">
          <div class="d-flex fill-height justify-center align-center">
            <v-img
                @click.stop="imgOverlay = false"
                max-width="85wh"
                max-height="90vh"
                height="auto"
                :src="replaceLargeImg(image)"
                :lazy-src="image.url"
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

  <v-snackbar
      location="top"
      color="success"
      v-model="snackbar"
      timeout="1000"
  >
    copy success
  </v-snackbar>
</template>

<style scoped>
.screen-height {
  height: calc(100vh - 80px);
}
</style>