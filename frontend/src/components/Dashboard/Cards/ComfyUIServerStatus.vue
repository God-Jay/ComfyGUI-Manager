<script setup>
import {ref} from "vue";
import {CheckIsServerRunning, StartServer, StopServer} from "@wailsjs/go/comfyUI/ComfyUI.js";
import BaseCard from "@/components/Dashboard/Cards/BaseCard.vue";
import {useMainStore} from "@/stores/store.js";

const serverStatusEnums = {
  Starting: 0,
  Running: 1,
  Stopping: 2,
  Stopped: 3,
}
const serverStatus = ref(serverStatusEnums.Stopped)

const snackbar = ref(false)
const timeout = 3000

const getIsServerRunning = () => {
  CheckIsServerRunning().then(r => {
    if (r) {
      serverStatus.value = serverStatusEnums.Running
      useMainStore().setServerIsRunning(true)
    } else {
      serverStatus.value = serverStatusEnums.Stopped
      useMainStore().setServerIsRunning(false)
    }
  })
}
getIsServerRunning()

const startServer = () => {
  serverStatus.value = serverStatusEnums.Starting
  StartServer().then(r => {
    getIsServerRunning()
  })
}

const stopServer = () => {
  serverStatus.value = serverStatusEnums.Stopping
  StopServer().then(r => {
    if (!r) {
      snackbar.value = true
    }
    getIsServerRunning()
  })
}

</script>

<template>
  <BaseCard>
    <v-card-title>ComfyUI Server Status</v-card-title>
    <v-row>
      <v-col cols="12" md="12">
        <v-card hover variant="tonal" color="primary" height="160" rounded="lg">
          <v-card-title>ComfyUI Server Info</v-card-title>
          <v-card-text v-if="true">
            <p>Port: 8188</p>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="12">
        <v-card hover variant="tonal" color="primary" height="160" rounded="lg">
          <v-card-title>Server Status</v-card-title>

          <template v-if="serverStatus === serverStatusEnums.Running">
            <v-card-text>
              <p class="text-success">
                <v-icon color="green" class="mr-1 mb-1">mdi-check-circle</v-icon>
                Server Now Running
              </p>
            </v-card-text>
            <v-card-actions style="position: absolute; bottom: 0">
              <v-btn border color="red" @click.stop="stopServer" class="ml-1">
                Stop Server
              </v-btn>
            </v-card-actions>
          </template>

          <template v-if="serverStatus === serverStatusEnums.Stopped">
            <v-card-text>
              <p class="text-red">
                <v-icon color="red" class="mr-1 mb-1">mdi-alert</v-icon>
                Server is Not Running
              </p>
            </v-card-text>
            <v-card-actions style="position: absolute; bottom: 0">
              <v-btn border color="success" @click.stop="startServer" class="ml-1">
                Start Server
              </v-btn>
            </v-card-actions>
          </template>

          <template v-if="serverStatus === serverStatusEnums.Starting">
            <v-card-text>
              <v-progress-circular
                  color="green"
                  indeterminate
                  class="mb-1"
              ></v-progress-circular>
              <span class="ml-4">Starting...</span>
            </v-card-text>
          </template>

          <template v-if="serverStatus === serverStatusEnums.Stopping">
            <v-card-text>
              <v-progress-circular
                  color="red"
                  indeterminate
                  class="mb-1"
              ></v-progress-circular>
              <span class="ml-4">Stopping...</span>
            </v-card-text>
          </template>

        </v-card>
      </v-col>

    </v-row>
  </BaseCard>

  <!--  stop server failed-->
  <v-snackbar
      location="top"
      color="red"
      v-model="snackbar"
      :timeout="timeout"
  >
    Cannot stop server, may be the server is running from the terminal.
    <template v-slot:actions>
      <v-btn
          variant="text"
          @click="snackbar = false"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<style scoped>

</style>