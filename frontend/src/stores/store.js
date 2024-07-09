import {defineStore} from 'pinia'

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useMainStore = defineStore('mainStore', {
    // a function that returns a fresh state
    state: () => ({
        comfyUIPath: '',
        mainNav: 'dashboard',
        welcome: true,
        serverIsRunning: false,
    }),
    // optional getters
    getters: {},
    // optional actions
    actions: {
        setMainNav(nav) {
            this.mainNav = nav
        },
        setWelcome(welcome) {
            this.welcome = welcome
        },
        setComfyUIPath(path) {
            this.comfyUIPath = path
        },
        setServerIsRunning(isRunning) {
            this.serverIsRunning = isRunning
        }
    },
})