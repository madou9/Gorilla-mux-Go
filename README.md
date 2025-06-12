# ProjectX – Employee Management API

A simple RESTful API built in Go using **Gorilla Mux** and **MongoDB**. It supports basic CRUD operations on employee data and follows a clean, modular structure (`usecase`, `repository`, `model`).

---

## Features

- Create a new employee
- Fetch employee by ID
- List all employees
- Update employee by ID
- Delete employee by ID
- Delete all employees
- Health check endpoint

---

## Tech Stack

- Language: **Go**
- Web Framework: **Gorilla Mux**
- Database: **MongoDB**
- UUID Generation: **google/uuid**
- Environment Config: **joho/godotenv**

---

## Project Structure

projectx/
├── main.go
├── model/
│ └── employee.go
├── repository/
│ └── employee.go
├── usecase/
│ └── employee.go
├── .env
└── go.mod
