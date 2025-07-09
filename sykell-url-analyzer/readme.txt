Sure! Here's a complete and professional README.md for your Sykell URL Analyzer full-stack project with React + TypeScript (frontend), Go (Gin), MySQL, and Docker Compose:

📄 README.md

# 🔍 Sykell URL Analyzer

A full-stack web application that crawls submitted URLs and displays key metadata including HTML version, headings, internal/external/broken links, and login form presence. Built for fast analysis with a TypeScript React frontend, a Go backend (Gin), and MySQL.

## ✨ Features

* ✅ Submit URLs to analyze web pages
* ✅ Real-time status updates (Queued → Running → Done)
* ✅ Dashboard with filtering, search, and sort
* ✅ View link stats and broken links chart
* ✅ Bulk re-run or delete analysis
* ✅ Authenticated API requests
* ✅ Responsive UI (mobile + desktop)

## 📦 Tech Stack

| Layer    | Technology                            |
| -------- | ------------------------------------- |
| Frontend | React, TypeScript, Tailwind, Recharts |
| Backend  | Go (Golang), Gin, GORM                |
| Database | MySQL                                 |
| Infra    | Docker, Docker Compose                |
| Auth     | Bearer token-based API auth           |
| Testing  | React Testing Library + Jest          |

## 🚀 Getting Started

### 1. Clone this repo

```bash
git clone https://github.com/yourusername/sykell-url-analyzer.git
cd sykell-url-analyzer
```

### 2. Run with Docker Compose

Make sure Docker is installed and running.

```bash
docker compose up --build
```

* Frontend: [http://localhost:3000](http://localhost:3000)
* Backend: [http://localhost:8080/api](http://localhost:8080/api)

### 3. Access MySQL (optional)

MySQL is exposed on port 3306 with:

* Username: root
* Password: password
* Database: sykell

You can connect with tools like DBeaver or TablePlus.

## 🧪 Running Frontend Tests

```bash
cd frontend
npm install
npm test
```

## 🗃 Folder Structure

```
sykell-url-analyzer/
├── backend/
│   ├── cmd/             # app entry point
│   ├── controllers/     # route handlers
│   ├── crawlers/        # HTML analyzer logic
│   ├── middleware/      # auth
│   ├── models/          # DB models
│   ├── routes/          # router setup
│   ├── utils/           # helper functions
│   ├── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   ├── services/
│   │   ├── hooks/
│   ├── Dockerfile
├── docker-compose.yml
├── README.md
```

## 🔐 Authorization

All backend API routes require a Bearer token in the Authorization header:

Example:

```http
GET /api/urls
Authorization: Bearer <your-token>
```

For demo purposes, token validation is mocked and can be easily extended.

## 📈 Example Output

* HTML Version: HTML5
* Title: Welcome to Example.com
* H1: 2
* Internal Links: 10
* External Links: 5
* Broken Links: 2 (404, 500)
* Login Form: Present

## 📌 To Do / Nice-to-Have

* [ ] OAuth2 or JWT auth backend
* [ ] WebSocket support for live crawl status
* [ ] Export report to PDF/CSV
* [ ] Admin dashboard for crawl limits/logs

## 📄 License

This project is for evaluation/demo purposes only. Feel free to reuse or open-source as needed.

---

Let me know if you want me to auto-fill your GitHub repo URL or adapt this for a private repo or a Heroku/Render/Netlify deployment.
