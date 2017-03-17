import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import Detail from '@/components/Detail'

var VueResource = require('vue-resource')

Vue.use(Router)
Vue.use(VueResource)
Vue.http.options.emulateJSON = true;
Vue.http.options.headers={
	'Content-Type':'application/x-www-form-urlencoded; charset=UTF-8'
};
Vue.http.options.xhr = { withCredentials: true }

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Hello',
      component: Hello
    },
    {
      path: '/detail/:id',
      name: 'Detail',
      component: Detail
    },
  ]
})
