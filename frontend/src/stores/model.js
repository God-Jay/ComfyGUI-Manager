import {defineStore} from 'pinia'

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useModelStore = defineStore('modelStore', {
    // a function that returns a fresh state
    state: () => ({
        modelDir: {},
        modelNav: 'main',
        specificModelDir: {},
        modelIndex: null
    }),
    // optional getters
    getters: {},
    // optional actions
    actions: {
        setModelDir(modelDir) {
            // `this` is the store instance
            this.modelDir = modelDir
        },
        setMainNav() {
            this.modelNav = 'main'
            this.modelIndex = null
        },
        setDirIndex(nav, index) {
            // `this` is the store instance
            this.specificModelDir = this.modelDir.subDirs[index]
            this.modelIndex = index
            this.modelNav = nav
        }
    },
})