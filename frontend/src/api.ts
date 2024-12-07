import type { Feature } from "./models"

export async function getFeatureList() {
    try {
        const resp = await fetch("/api/feature", {
            method: "GET"
        })
        const body = await resp.json()
        return body
    } catch (err) {
        console.error(err)
    }
}

export async function saveFeature(feature: Feature) {
    try {
        const resp = await fetch("/api/feature", {
            method: "POST",
            body: JSON.stringify(feature),
            headers: {
                "Content-Type": "application/json",
            },
        })
        const body = await resp.json()
        feature.ID = body
    } catch (err) {
        console.error(err)
    }
}

export async function updateFeature(feature: Feature) {
    try {
        const resp = await fetch("/api/feature", {
            method: "PUT",
            body: JSON.stringify(feature),
            headers: {
                "Content-Type": "application/json",
            },
        })
    } catch (err) {
        console.error(err)
    }
}

export async function deleteFeature(id: number) {
    try {
        await fetch(`/api/feature/${id}` , {
            method: "DELETE",
        })
    } catch (err) {
        console.error(err)
    }
}