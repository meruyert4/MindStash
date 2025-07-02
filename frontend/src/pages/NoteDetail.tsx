import { useNavigate, useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import { deleteNote, getNote } from "../api/notes";

export default function NoteDetail() {
  const { id } = useParams();
  const [note, setNote] = useState<any>(null);
  const navigate = useNavigate();

  useEffect(() => {
    if (id) {
      getNote(id).then(setNote);
    }
  }, [id]);

  const handleDelete = async () => {
    if (id) {
      await deleteNote(id);
      navigate("/");
    }
  };

  if (!note) return <div className="p-4">Loading...</div>;

  return (
    <div className="max-w-2xl mx-auto p-4">
      <h1 className="text-3xl font-bold mb-2">{note.title}</h1>
      <p className="text-gray-700 mb-4">{note.content}</p>
      <div className="flex gap-4">
        <button
          onClick={() => navigate(`/notes/${id}/edit`)}
          className="bg-yellow-500 text-white px-4 py-2 rounded hover:bg-yellow-600"
        >
          Edit
        </button>
        <button
          onClick={handleDelete}
          className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700"
        >
          Delete
        </button>
      </div>
    </div>
  );
}
