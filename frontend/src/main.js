import Vue from 'vue'
import App from './App.vue'
import axios from 'axios'
import ElementUI from 'element-ui' //element-ui的全部组件
import 'element-ui/lib/theme-chalk/index.css'//element-ui的css
Vue.use(ElementUI);//使用elementUI

Vue.config.productionTip = false;
Vue.prototype.$ajax = axios;
new Vue({
  render: h => h(App),
}).$mount('#app');


