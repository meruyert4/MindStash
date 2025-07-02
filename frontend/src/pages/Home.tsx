import { useEffect, useState } from "react";
import { getNotes } from "../api/notes";
import NoteCard from "../components/NoteCard";
import { Link } from "react-router-dom";

interface Note {
  id: string;
  title: string;
  content: string;
}

export default function Home() {
  const [notes, setNotes] = useState<Note[]>([]);

  useEffect(() => {
    getNotes().then(setNotes);
  }, []);

  return (
    <div className="max-w-4xl mx-auto p-4">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-3xl font-bold">MindStash</h1>
        <Link to="/create" className="px-4 py-2 bg-indigo-600 text-white rounded hover:bg-indigo-700">
          + New Note
        </Link>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {notes.map(note => (
          <NoteCard key={note.id} {...note} />
        ))}
      </div>
    </div>
  );
}
