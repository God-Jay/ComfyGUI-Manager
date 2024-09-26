import {createApp} from 'vue'

// Vuetify
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader
import {createVuetify} from 'vuetify'

import {createPinia} from 'pinia'

// Components
import App from './App.vue'

const vuetify = createVuetify()
const pinia = createPinia()

createApp(App).use(vuetify).use(pinia).mount('#app')
