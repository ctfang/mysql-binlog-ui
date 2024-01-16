import "./assets/main.css";

import { createApp, ref, provide } from "vue";
import App from "./App.vue";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import router from "./router";

const app = createApp(App);

app.use(router);
app.use(ElementPlus);

app.mount("#app");
