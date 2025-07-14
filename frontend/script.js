const API_BASE = "http://localhost:8080";

// DOM Elements
const form = document.getElementById("note-form");
const notesList = document.getElementById("notes-list");
const themeToggle = document.getElementById("theme-toggle");
const body = document.body;
const notesCount = document.getElementById("notes-count");

// Theme management
const currentTheme = localStorage.getItem('theme') || 'light';
setTheme(currentTheme);

themeToggle.addEventListener('click', () => {
  const newTheme = body.classList.contains('dark-theme') ? 'light' : 'dark';
  setTheme(newTheme);
  localStorage.setItem('theme', newTheme);
});

function setTheme(theme) {
  if (theme === 'dark') {
    body.classList.add('dark-theme');
    themeToggle.innerHTML = `
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M21 12.79C20.8427 14.4922 20.2039 16.1144 19.1582 17.4668C18.1125 18.8192 16.7035 19.8458 15.0957 20.4265C13.4879 21.0073 11.748 21.1181 10.0795 20.7461C8.41102 20.3741 6.8829 19.5345 5.67418 18.3258C4.46545 17.1171 3.62593 15.589 3.2539 13.9205C2.88187 12.252 2.99268 10.5121 3.57345 8.9043C4.15422 7.29651 5.18082 5.88752 6.5332 4.84182C7.88558 3.79611 9.5078 3.15732 11.21 3C10.2134 4.34827 9.73375 6.00945 9.85849 7.68141C9.98324 9.35338 10.7039 10.9251 11.8894 12.1106C13.0749 13.2961 14.6466 14.0168 16.3186 14.1415C17.9906 14.2663 19.6517 13.7866 21 12.79Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>
    `;
  } else {
    body.classList.remove('dark-theme');
    themeToggle.innerHTML = `
      <svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M12 18C15.3137 18 18 15.3137 18 12C18 8.68629 15.3137 6 12 6C8.68629 6 6 8.68629 6 12C6 15.3137 8.68629 18 12 18Z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M12 2V4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M12 20V22" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M4.92993 4.92999L6.33993 6.33999" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M17.6599 17.66L19.0699 19.07" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M2 12H4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M20 12H22" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M4.92993 19.07L6.33993 17.66" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
        <path d="M17.6599 6.33999L19.0699 4.92999" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
      </svg>
    `;
  }
}

// Note CRUD operations
form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const title = document.getElementById("title").value;
  const content = document.getElementById("content").value;

  try {
    const res = await fetch(`${API_BASE}/notes`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title, content }),
    });

    if (!res.ok) throw new Error("Failed to create note");

    form.reset();
    loadNotes();
  } catch (err) {
    alert("Error creating note!");
    console.error(err);
  }
});

async function loadNotes() {
  notesList.innerHTML = "";
  try {
    const res = await fetch(`${API_BASE}/notes`);
    if (!res.ok) throw new Error("Failed to fetch notes");

    const notes = await res.json();
    updateNotesCount(notes.length);

    if (notes.length === 0) {
      notesList.innerHTML = "<p class='empty-message'>🪶 Нет заметок пока что!</p>";
      return;
    }

    notes.reverse().forEach((note) => {
      const el = document.createElement("div");
      el.className = "note";
      el.innerHTML = `
        <h3>${note.title}</h3>
        <p>${note.content}</p>
        <div class="note-actions">
          <button class="delete-btn" data-id="${note.id}">Delete</button>
        </div>
      `;
      notesList.appendChild(el);

      el.querySelector('.delete-btn').addEventListener('click', async () => {
        try {
          const deleteRes = await fetch(`${API_BASE}/notes/${note.id}`, {
            method: "DELETE"
          });
          if (!deleteRes.ok) throw new Error("Failed to delete note");
          loadNotes();
        } catch (err) {
          alert("Error deleting note!");
          console.error(err);
        }
      });
    });

  } catch (err) {
    console.error("Error loading notes:", err);
    notesList.innerHTML = "<p class='error-message'>Notes Not Found. Create One!</p>";
  }
}

function updateNotesCount(count) {
  notesCount.textContent = `${count} ${count === 1 ? 'note' : 'notes'}`;
}

// Initial load
loadNotes();