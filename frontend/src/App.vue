<template>
  <v-app>
    <v-layout class="rounded rounded-md h-100">
      <v-navigation-drawer location="right" width="400">
        <v-list>
          <v-list-item title="Feature Planner"></v-list-item>
          <v-list-item>
            <v-btn block variant="outlined" @click="openNewForm">
              <template #prepend>
                <v-icon icon="mdi-plus-thick"></v-icon>
              </template> Create new feature
            </v-btn>
          </v-list-item>
          <v-list-item v-if="showForm">
            <h3 class="pb-2">Feature-Data</h3>
            <feature-form @clear="clear" @save="save"></feature-form>
          </v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-main class="d-flex align-center justify-center flex-column ga-6">
        <transition-group name="list">
          <feature-card v-for="(feature) in store.sorted" :key="feature.ID" :feature="feature" @edit="edit"
            @remove="remove" />
        </transition-group>
      </v-main>
    </v-layout>
  </v-app>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { featureSkel } from './models'
import type { Feature } from './models'
import { useFeaturesStore } from './store'
import FeatureCard from './components/FeatureCard.vue';
import FeatureForm from './components/FeatureForm.vue';

const store = useFeaturesStore()
store.init()
const showForm = ref(false)

function openNewForm() {
  store.editFeature = featureSkel()
  showForm.value = true
}

function clear() {
  store.editFeature = featureSkel()
  showForm.value = false
}

function save() {
  const feature = JSON.parse(JSON.stringify(store.editFeature))
  store.add(feature)
  clear()
}

function edit(feature: Feature) {
  store.editFeature = JSON.parse(JSON.stringify(feature))
  showForm.value = true
}

function remove(feature: Feature) {
  store.remove(feature)
}

</script>

<style>
.list-move, /* apply transition to moving elements */
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}

.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

/* ensure leaving items are taken out of layout flow so that moving
   animations can be calculated correctly. */
.list-leave-active {
  position: absolute;
}
</style>