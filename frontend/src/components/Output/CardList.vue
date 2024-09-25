<script setup>
import {onActivated, ref} from "vue";
import {OpenFileInDir} from "@wailsjs/go/backend/App.js";
import moment from "moment";

defineProps({
  images: {
    type: Array,
    required: true
  },
  clickImg: {
    type: Function,
    required: true
  },
  setImgRef: {
    type: Function,
    required: true
  },
  starImg: {
    type: Function,
    required: true
  },
  viewWorkflow: {
    type: Function,
    required: true
  }
})

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
                 class="overflow-y-auto screen-height-tab">
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
                :ref="setImgRef(i, imageFile)"
            >
              <v-toolbar color="transparent">
                <template v-slot:append>
                  <v-btn variant="tonal" icon="mdi-star" :color="imageFile.stared ? 'red': 'white'"
                         @click.stop="starImg(imageFile)"></v-btn>
                </template>
              </v-toolbar>
            </v-img>
            <v-card-text class="text-center pt-1 pb-1">
              <v-btn color="primary" size="small" class="mb-2" @click="OpenFileInDir('output', imageFile.name)">
                open in path
              </v-btn>
              <v-btn color="success" size="small" class="mb-2" @click.stop="viewWorkflow(imageFile, i)">
                view workflow
              </v-btn>
              <p>{{ imageFile.name }}</p>
              <p>{{ imageFile.ref?.naturalWidth + '×' + imageFile.ref?.naturalHeight + ' pix' }}</p>
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
</template>

<style scoped>
.screen-height-tab {
  height: calc(100vh - 80px - 48px);
}
</style>