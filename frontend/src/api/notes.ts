const API_BASE = "http://backend:8080";

export async function getNotes() {
  const res = await fetch(`${API_BASE}/notes`);
  return await res.json();
}

export async function getNote(id: string) {
  const res = await fetch(`${API_BASE}/notes/${id}`);
  return await res.json();
}

export async function createNote(note: { title: string; content: string }) {
  const res = await fetch(`${API_BASE}/notes`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(note),
  });

  if (!res.ok) {
    const text = await res.text();
    console.error("Backend error:", text);
    throw new Error("Failed to create note");
  }

  // <-- эта строка вызывала ошибку, если тела нет
  if (res.status === 204 || res.headers.get("content-length") === "0") {
    return null; // ничего не возвращаем
  }

  return await res.json(); // вернём JSON, если он есть
}

export async function updateNote(id: string, data: { title: string; content: string }) {
  const res = await fetch(`${API_BASE}/notes/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  return await res.json();
}

export async function deleteNote(id: string) {
  return await fetch(`${API_BASE}/notes/${id}`, {
    method: "DELETE",
  });
}
