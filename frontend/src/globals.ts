export function estimateColor(estimate: string): string {
  switch (estimate) {
    case "short":
      return "light-blue-lighten-2"
    case "medium":
      return "orange-lighten-2"
    case "long":
      return "red-lighten-2"
    default:
      return "primary"
  }
}

export function featureColor(type: string): string {
  switch (type) {
    case "Bug":
      return "red-lighten-2"
    case "Feature":
      return "orange-lighten-2"
    case "Improvement":
      return "light-blue-lighten-2"
    default:
      return "light-blue-lighten-2"
  }
}

