import Vue from 'vue'
import App from './components/App'
import router from './router'
import './services'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
	el: '#app',
	router,
	template: '<App/>',
	components: { App }
})
