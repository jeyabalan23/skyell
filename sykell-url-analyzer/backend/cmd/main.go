
// backend/cmd/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "sykell-url-analyzer/routes"
    "sykell-url-analyzer/database"
)

func main() {
    db := database.InitDB()
    r := gin.Default()
    routes.SetupRoutes(r, db)
    r.Run(":8080")
}

// frontend/src/pages/Home.tsx
import URLForm from "../components/URLForm"
import DashboardTable from "../components/DashboardTable"

const Home = () => (
  <div className="p-6 space-y-6">
    <URLForm />
    <DashboardTable />
  </div>
)

export default Home

// frontend/src/services/api.ts
import axios from "axios"

export const api = axios.create({
  baseURL: "http://localhost:8080/api",
  headers: {
    Authorization: "Bearer mysecrettoken",
  },
})
