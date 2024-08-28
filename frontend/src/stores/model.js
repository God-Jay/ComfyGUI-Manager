import {defineStore} from 'pinia'

// main is the name of the store. It is unique across your application
// and will appear in devtools
export const useModelStore = defineStore('modelStore', {
    // a function that returns a fresh state
    state: () => ({
        modelDir: {},

        // empty to be model main path
        selectedModelPath: '',
        selectedModelDir: {}
    }),
    // optional getters
    getters: {},
    // optional actions
    actions: {
        setModelDir(modelDir) {
            // `this` is the store instance
            this.modelDir = modelDir
        },

        isMainPath() {
            return this.selectedModelPath === ''
        },
        selectModelPath(path) {
            this.selectedModelPath = path
            this.selectedModelDir = this.findDir(this.modelDir.subDirs, path)
        },
        findDir(dirs, path) {
            const dirNames = path.split('/');

            let currentDirs = dirs;
            let result = null;

            for (const dirName of dirNames) {
                result = currentDirs.find(dir => dir.name === dirName);
                if (!result) {
                    return null;
                }
                currentDirs = result.subDirs;
            }

            return result;
        }
    },


})

