<template>
  <v-list>
    <v-list-item title="Features"></v-list-item>
    <v-list-item>
      <v-btn block variant="outlined" @click="openNewForm">
        <template #prepend>
          <v-icon icon="mdi-plus-thick"></v-icon>
        </template>
        Create new feature
      </v-btn>
    </v-list-item>
    <v-list-item v-if="store.showForm">
      <h3 class="pb-2">Feature-Data</h3>
      <feature-form @clear="clear" @save="save"></feature-form>
    </v-list-item>
  </v-list>
</template>

<script setup lang="ts">
import FeatureForm from "@/components/FeatureForm.vue";
import {featureSkel} from "@/models";
import {useFeaturesStore} from "@/store/feature";

const store = useFeaturesStore()

function openNewForm() {
  store.editFeature = featureSkel()
  store.showForm = true
}

function clear() {
  store.editFeature = featureSkel()
  store.showForm = false
}

function save() {
  const feature = JSON.parse(JSON.stringify(store.editFeature))
  store.add(feature)
  clear()
}

</script>
