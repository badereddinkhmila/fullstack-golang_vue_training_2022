import { createApp, provide, h } from 'vue';
import App from './App.vue';
import router from './router';
import ElementPlus from 'element-plus';
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import { createPinia } from 'pinia';
import piniaPersist from 'pinia-plugin-persist';
import { DefaultApolloClient } from '@vue/apollo-composable';
import { apolloClient } from './api/apolloClient';

import './styles/index.scss';
import './element-variables.scss';

const app: any = createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient);
  },
  render: () => h(App),
});

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

const pinia = createPinia();
pinia.use(piniaPersist);

app
  .provide(DefaultApolloClient, apolloClient)
  .use(ElementPlus)
  .use(pinia)
  .use(router)
  .mount('#app');
