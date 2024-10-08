import {defineStore} from 'pinia'

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useMainStore = defineStore('mainStore', {
    // a function that returns a fresh state
    state: () => ({
        comfyUIPath: '',
        mainNav: 'dashboard',
        serverIsRunning: false,

        currentPath: {path: '', fileName: ''},
    }),
    // optional getters
    getters: {
        hasSetComfyUIPath() {
            console.log(this.comfyUIPath)
            return this.comfyUIPath !== ''
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
        setCurrentPath(path, fileName) {
            console.log("setCurrentPath", path, fileName)
            // `this` is the store instance
            this.currentPath.path = path
            this.currentPath.fileName = fileName
        },
        setServerIsRunning(isRunning) {
            this.serverIsRunning = isRunning
        }
    },
})