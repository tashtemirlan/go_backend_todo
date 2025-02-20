# 🚀 Material To-Do Backend

This project is a **Go (Gin) backend** for a **To-Do application** using **PostgreSQL** as the database.

## 🎯 Getting Started

### 🔍 Check PostgreSQL Status
To check status of PostgreSQL is working please run :
```
pg_ctl status -D <path to postgre>\PostgreSQL\data
```
### ▶️ Start PostgreSQL
To start work PostgreSQL please run :
```
pg_ctl start -D <path to postgre>\PostgreSQL\data
```
### 🛠️ Create the Database
Log in to your postgres sql database and create table by running this command :
```
CREATE DATABASE material_todo_go;
```

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