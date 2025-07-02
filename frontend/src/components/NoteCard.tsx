import { Link } from "react-router-dom";

interface NoteCardProps {
  id: string;
  title: string;
  content: string;
}

export default function NoteCard({ id, title, content }: NoteCardProps) {
  return (
    <Link to={`/notes/${id}`} className="block p-4 border rounded-lg bg-white shadow hover:shadow-md transition">
      <h2 className="text-xl font-semibold">{title}</h2>
      <p className="text-gray-600 mt-2">{content.slice(0, 100)}...</p>
    </Link>
  );
}
