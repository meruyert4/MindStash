# 🧠 MindStash

A minimalistic full-stack note-taking app (like Notion/Google Keep) built with:

- **Backend:** Go + PostgreSQL + Redis + MinIO
- **Frontend:** React + TypeScript + Tailwind + Vite
- **DevOps:** Docker Compose
---

## 🚀 Features

- 📝 Create, edit, delete notes
- 📂 Notes stored in PostgreSQL
- 🧪 API tested via Postman collection

---

## 📦 Tech Stack

| Layer      | Stack                          |
|------------|--------------------------------|
| Backend    | Go (net/http) + PostgreSQL     |
| Frontend   | React + TypeScript + Tailwind  |
| Dev Tools  | Docker Compose, Postman        |

---

## 🛠️ Setup & Run Locally

### 1. 🔃 Clone the repo

```bash
git clone https://github.com/meruyert4/mindstash.git
cd mindstash
````

---

### 2. 🐳 Start the full stack with Docker

```bash
docker-compose up --build
```

> This will start:
>
> * Go backend on [http://localhost:8080](http://localhost:8080)
> * React frontend on [http://localhost:5173](http://localhost:5173)
> * PostgreSQL at port `5434`

---

### 3. 🧪 API Testing (Optional)

You can test the backend APIs using **Postman**.

* 💼 Collection available at:
  [`/tests/MindStash.postman_collection.json`](./tests/MindStash.postman_collection.json)

Just import it into Postman and test all the endpoints:

* `GET /notes`
* `POST /notes`
* `GET /notes/{id}`
* `PUT /notes/{id}`
* `DELETE /notes/{id}`

---

## 🧍 Author

Built with 💜 by [@meruyert4](https://github.com/meruyert4)

---