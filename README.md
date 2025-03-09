# 🚀 Material To-Do Backend

This project is a **Go (Gin) backend** for a **To-Do application** using **PostgreSQL** as the database.

## 🎯 Getting Started

### 🏭 Project Structure
```
📂 material_todo_go
 ├── 📁 controllers    # API controllers (auth, user, notes)
 ├── 📁 models         # Database models (User, Note)
 ├── 📁 database       # Database connection (PostgreSQL)
 ├── 📁 routes         # API route definitions
 ├── 📁 utils          # Utility functions (JWT handling)
 ├── 📁 uploads        # Folder for user-uploaded images
 ├── 🌍 main.go        # Entry point of the app
 ├── 📜 go.mod         # Go modules
 ├── 📜 README.md      # Project documentation
```

### 🛠️ Running the Project
```sh
docker-compose up --build -d
```