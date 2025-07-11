
// frontend/src/services/api.ts
import axios from "axios"

export const api = axios.create({
  baseURL: "http://localhost:8080/api",
  headers: {
    Authorization: "Bearer mysecrettoken",
  },
})
