<script setup>
import {ref} from "vue";
import {GetImages} from "@wailsjs/go/output/Output.js";
import moment from "moment";

const images = ref()

const getImages = () => {
  GetImages().then(r => {
    images.value = r
  })
}
getImages()
</script>

<template>

  <v-card
      elevation="0"
      class="mx-auto"
      min-width="860"
  >
    <v-container fluid>
      <v-row dense>
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
                :src="'http://localhost:8190/output/' + imageFile.name"
                cover
            ></v-img>
            <v-card-text class="text-center">
              <p>{{ imageFile.name }}</p>
              <p>{{ imageFile.size }}</p>
              <p>{{ moment(imageFile.modTime).format('LLL') }}</p>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-card>

</template>

<style scoped>

</style>