import { createApp } from "vue";
import "./assets/main.css";
import router from "./router";
import App from "./App.vue";
import "primeicons/primeicons.css";
import Toast from "vue3-toastify";
import "vue3-toastify/dist/index.css";

const app = createApp(App);

app.use(Toast);

app.use(router);

app.mount("#app");
