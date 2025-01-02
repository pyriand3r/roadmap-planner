import { createMemoryHistory, createRouter } from 'vue-router'

import Features from "@/views/Features.vue"
import Roadmap from "@/views/Roadmap.vue"

const routes = [
  { path: '/', redirect: '/features' },
  { path: '/features', component: Features },
  { path: '/roadmap', component: Roadmap },
]

export default createRouter({
  history: createMemoryHistory(),
  routes,
})
