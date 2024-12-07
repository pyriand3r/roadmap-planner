
export const feature_type_list = [
  "Feature",
  "Improvement",
  "Bug",
]

export type FeatureType = "" | "Feature" | "Improvement" | "Bug"

export type Feature = {
  ID: number,
  title: string,
  description: string,
  customer_score: number,
  internal_score: number,
  sales_score: number,
  impact_score: number,
  total_score: number,
  type: FeatureType
}

export function featureSkel(): Feature {
  return {
    ID: 0,
    title: "",
    description: "",
    customer_score: 0,
    internal_score: 0,
    sales_score: 0,
    impact_score: 0,
    total_score: 0,
    type: ""
  }
}