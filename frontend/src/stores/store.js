import {defineStore} from 'pinia'

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useMainStore = defineStore('mainStore', {
    // a function that returns a fresh state
    state: () => ({
        comfyUIPath: '',
        mainNav: 'dashboard',
        serverIsRunning: false,
    }),
    // optional getters
    getters: {
        hasSetComfyUIPath() {
            console.log(this.comfyUIPath)
            console.log(this.comfyUIPath !== '')
            return this.comfyUIPath !== ''
        },
        getAA() {
            return this.comfyUIPath
        }
    },
    // optional actions
    actions: {
        setMainNav(nav) {
            this.mainNav = nav
        },
        setComfyUIPath(path) {
            this.comfyUIPath = path
        },
        setServerIsRunning(isRunning) {
            this.serverIsRunning = isRunning
        }
    },
})