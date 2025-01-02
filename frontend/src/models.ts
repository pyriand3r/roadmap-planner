
export const feature_type_list = [
  "Feature",
  "Improvement",
  "Bug",
]

export const feature_estimate_list = [
  "short",
  "medium",
  "long",
]

export type FeatureEstimate = "short" | "medium" | "long"

export type FeatureType = "" | "Feature" | "Improvement" | "Bug"

export type Feature = {
  ID: number,
  title: string,
  description: string,
  customer_score: number,
  dev_score: number,
  sales_score: number,
  total_score: number,
  type: FeatureType,
  estimate: FeatureEstimate,
  milestone_id: number,
}

export function featureSkel(): Feature {
  return {
    ID: 0,
    title: "",
    description: "",
    customer_score: 0,
    dev_score: 0,
    sales_score: 0,
    total_score: 0,
    type: "",
    estimate: "medium",
    milestone_id: 0,
  }
}

export type Milestone = {
  ID: number,
  title: string,
  description: string,
  features: number[],
}

export function milestoneSkel(): Milestone {
  return {
    ID: 0,
    title: "",
    description: "",
    features: [],
  }
}