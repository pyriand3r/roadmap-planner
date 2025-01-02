<template>
  <v-app>
    <v-layout class="rounded rounded-md h-100 w-100 position-fixed">
      <v-app-bar title="Roadmap Planner" elevation="2">
        <v-tooltip text="Roadmap" location="bottom">
          <template #activator="{ props }">
            <v-btn v-bind="props" icon="mdi-sign-direction" to="/roadmap" />
          </template>
        </v-tooltip>
        <v-tooltip text="Features" location="bottom">
          <template #activator="{ props }">
            <v-btn v-bind="props" icon="mdi-star" to="/features" />
          </template>
        </v-tooltip>
        <v-app-bar-nav-icon variant="text" @click.stop="store.sidebarShow = !store.sidebarShow" />
      </v-app-bar>
      <RouterView />
    </v-layout>
  </v-app>
</template>

<script setup lang="ts">
import Features from '@/views/Features.vue'
import { useFeaturesStore } from "@/store/feature"
import { useTheme } from 'vuetify'
const theme = useTheme()

if (window.matchMedia) {
  if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
    theme.global.name.value = "dark"
}
  window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', event => {
    const newColorScheme = event.matches ? "dark" : "light"
    theme.global.name.value = newColorScheme 
  })
}

const store = useFeaturesStore()
</script>
