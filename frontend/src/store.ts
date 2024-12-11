import {defineStore} from 'pinia'
import {computed, ref} from 'vue'
import {featureSkel, type Feature} from './models'
import {deleteFeature, getFeatureList, saveFeature, updateFeature} from './api'

export const useFeaturesStore = defineStore('features', () => {
  const features = ref<Feature[]>([])
  const editFeature = ref<Feature>(featureSkel())
  const showForm = ref(false)

  async function add(feature: Feature) {
    calculateTotal(feature)
    if (!feature.ID || feature.ID === 0) {
      await saveFeature(feature)
      features.value.push(feature)
      return
    }
    await updateFeature(feature)
    features.value.forEach((val, idx) => {
      if (val.ID === feature.ID) {
        features.value[idx] = feature
      }
    })
  }

  async function remove(feature: Feature) {
    await deleteFeature(feature.ID)
    for (let i = 0; i < features.value.length; i += 1) {
      if (features.value[i].ID === feature.ID) {
        features.value.splice(i, 1)
        break
      }
    }
  }

  const sorted = computed(() => {
    if (features.value.length <= 1) {
      return features.value
    }
    const sorted = features.value.sort((a, b) => b.total_score - a.total_score)
    return sorted
  })

  async function init() {
    features.value = await getFeatureList()
    features.value.forEach((val) => {
      calculateTotal(val)
    })

  }

  function calculateTotal(feature: Feature) {
    feature.total_score = 2 * feature.customer_score + feature.internal_score + feature.impact_score + feature.sales_score
  }

  return {features, editFeature, add, sorted, init, remove}
})
