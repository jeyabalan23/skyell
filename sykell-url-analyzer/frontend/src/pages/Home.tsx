
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
