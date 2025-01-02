import {defineStore} from 'pinia'
import {ref} from 'vue'
import {milestoneSkel, type Milestone} from '../models'
import { getMilestoneList, saveMilestone, updateMilestone, deleteMilestone } from '../api'

export const useMilestoneStore = defineStore('milestones', () => {
  const sidebarShow = ref(true)
  const showForm = ref(false)
  let initialized = false

  const milestones = ref<Milestone[]>([])
  const editMilestone = ref<Milestone>(milestoneSkel())

  async function init() {
    if (initialized) {
      return
    }
    milestones.value = await getMilestoneList()
    initialized = true
  }

  async function add(milestone: Milestone) {
    if (!milestone.ID || milestone.ID === 0) {
      await saveMilestone(milestone)
      milestones.value.push(milestone)
      return
    }

    await updateMilestone(milestone)
    milestones.value.forEach((val, idx) => {
      if (val.ID === milestone.ID) {
        milestones.value[idx] = milestone
      }
    })
  }

  async function remove(milestone:Milestone) {
    await deleteMilestone(milestone.ID)
    for (let i = 0; i < milestones.value.length; i += 1) {
      if (milestone.ID === milestones.value[i].ID) {
        milestones.value.splice(i, 1)
        break
      }
    }
  }

  function byID(id: number): Milestone|null {
    for (let i = 0; i < milestones.value.length; i += 1) {
      if (milestones.value[i].ID === id) {
        return milestones.value[i]
      }
    }
    return null
  }

  return {init, sidebarShow, showForm, milestones, editMilestone, add, remove, byID}
})