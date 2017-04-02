import Vue from 'vue'
import Resource from 'vue-resource'

Vue.use(Resource)

Vue.http.options.root = 'http://localhost:8080/api/v0'
