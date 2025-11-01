# 🧠 go-microservices

Entire codebase is meant for **Golang practice** and for **building microservices using Go**.  
This repository serves as a hands-on lab for learning clean architecture, modular development, and API design using Go frameworks like **Chi** and **Gin**.

---

## 🎯 Objective

To understand and implement microservice architecture in Go —  
including routing, middleware, Dockerization, and API documentation using Swagger.

---

## 📺 YouTube Playlist

Learn along with the tutorial series I’m following:  
🎥 [Go Microservices by Nic Jackson](https://youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_&si=WErbZ-ktbhVNfJMV)

---

## 🧩 Key Components

- ⚙️ **Chi-based Services** — Lightweight and idiomatic Go microservices using the Chi router.  
- ⚡ **Gin-based Services** — High-performance REST APIs exploring Gin’s speed and middleware design.  
- 🧱 **Core Patterns** — Configuration management, structured logging, dependency injection.  
- 🧪 **Swagger Integration** — Auto-generated OpenAPI documentation for all routes.  
- 🐳 **Docker Support** — Dockerfiles for containerized builds and environment setup.  

---

## 🧰 Tech Stack

`Go` · `Chi` · `Gin` · `Swagger` · `Docker` · `Postman` · `JSON` · `Linux`

---

## Swagger Execution for reference:
1) go install github.com/go-swagger/go-swagger/cmd/swagger@latest => This gets all the binaries required for swagger
2) swagger generate spec -o ./swagger.yaml --scan-models => Generates the swagger.yaml file
3) I have a shortcut created generate-swagger.bat ! Execute: ./generate-swagger to recreate swagger.yaml file
4) Run your main function and Check your localhost:<"port no">/docs for the documentation
