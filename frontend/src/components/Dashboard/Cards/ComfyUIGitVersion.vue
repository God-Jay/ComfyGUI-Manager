<script setup>
import {GitHeadCommit, GitLatestCommit, GitPull, GitStatus} from "@wailsjs/go/comfyUI/ComfyUI.js";
import {ref} from "vue";
import moment from "moment";
import BaseCard from "@/components/Dashboard/Cards/BaseCard.vue";
import {EventsOn} from "@wailsjs/runtime/runtime.js";

const currentCommit = ref()
const getHeadCommit = () => {
  console.log("get head commit")
  GitHeadCommit().then(r => {
    currentCommit.value = r
    console.log(r)
  })
}

getHeadCommit()

const latestCommit = ref()
GitLatestCommit()
EventsOn("GitLatestCommit", (commit) => {
  latestCommit.value = commit
  if (commit) {
    getGitStatus()
  }
})

const hasCheck = ref(false)
const behindCommits = ref([])
const getGitStatus = () => {
  if (!hasCheck.value) {
    GitStatus().then(r => {
      behindCommits.value = r
      hasCheck.value = true
    })
  }
}
getGitStatus()

const refreshCheck = () => {
  hasCheck.value = false
  latestCommit.value = undefined
  GitLatestCommit()
}

const updating = ref(false)

const snackbar = ref(false)
const snackbarText = ref("Update Success")
const timeout = 3000

async function update() {
  updating.value = true
  const result = await GitPull()
  updating.value = false
  console.log("update result", result)
  snackbar.value = true
  snackbarText.value = "Update Success"
  getHeadCommit()
  getGitStatus()
}
</script>

<template>
  <BaseCard>
    <v-card-title>ComfyUI Git Version</v-card-title>
    <v-row>
      <v-col cols="12" md="6">
        <v-card hover variant="tonal" color="primary" height="160" rounded="lg">
          <v-card-title>Current ComfyUI Version</v-card-title>
          <v-card-text v-if="currentCommit">
            <p>Author: {{ currentCommit?.Author?.Name }}</p>
            <p>Updated At: {{ moment(currentCommit?.Author?.When).format('LLL') }}</p>
          </v-card-text>
        </v-card>
      </v-col>

      <v-col cols="12" md="6">
        <v-card hover variant="tonal" color="primary" height="160" rounded="lg">
          <v-card-title>Latest Update</v-card-title>
          <template v-if="updating">

            <v-card-text>
              <v-progress-circular
                  color="green"
                  indeterminate
              ></v-progress-circular>
              <span class="ml-4">Updating...</span>
            </v-card-text>

          </template>
          <template v-else>

            <template v-if="latestCommit">
              <v-card-text>
                <p>Author: {{ latestCommit?.Author?.Name }}</p>
                <p>Updated At: {{ moment(latestCommit?.Author?.When).format('LLL') }}</p>
                <template v-if="behindCommits.length === 0">
                  <p class="text-green">All files are up to date.</p>
                  <v-btn border color="primary" class="ml-1 mt-1" @click.stop="refreshCheck">Check</v-btn>
                </template>
                <p v-else class="text-red">{{ behindCommits.length }} commits behind head</p>
              </v-card-text>
              <v-card-actions style="position: absolute; bottom: 0"
                              v-if="behindCommits.length > 0">
                <v-spacer></v-spacer>
                <v-btn border @click.stop="update" color="success" class="ml-1">
                  Update
                </v-btn>
              </v-card-actions>
            </template>

            <template v-else>
              <v-card-text>
                <v-progress-circular
                    color="green"
                    indeterminate
                ></v-progress-circular>
                <span class="ml-4">Checking...</span>
              </v-card-text>
            </template>

          </template>
        </v-card>
      </v-col>

      <v-col cols="12" md="12">
        <v-card hover variant="tonal" color="primary" height="200" rounded="lg">
          <v-card-title>Recent Update Changes</v-card-title>
          <v-card-text v-if="behindCommits.length !== 0">
            <v-list height="140">
              <template v-for="commit in behindCommits">
                <v-list-item link color="red">
                  <v-list-item-title>{{ commit.Author.Name }}</v-list-item-title>
                  <v-list-item-subtitle>{{ commit.Message }}</v-list-item-subtitle>
                  <v-list-item-subtitle>{{ moment(commit.Author.When).format('LLL') }}</v-list-item-subtitle>
                  <v-divider></v-divider>
                </v-list-item>
              </template>
            </v-list>
          </v-card-text>
          <v-card-text v-else>No updates</v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </BaseCard>

  <!--  updating-->
  <v-snackbar
      location="top"
      color="primary"
      v-model="updating"
      timeout="-1"
  >
    <v-progress-circular
        color="green"
        indeterminate
    ></v-progress-circular>
    <span class="ml-4">Updating...</span>
  </v-snackbar>

  <!--  update result-->
  <v-snackbar
      close-on-content-click
      location="top"
      color="success"
      v-model="snackbar"
      :timeout="timeout"
  >
    {{ snackbarText }}
    <template v-slot:actions>
      <v-btn color="white" variant="text" @click="snackbar = false">
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<style scoped>

</style>