import { Route, Routes, Navigate } from "react-router-dom";
import Home from "./pages/Home";
import CreateNote from "./pages/CreateNote";
import NoteDetail from "./pages/NoteDetail";
import EditNote from "./pages/EditNote";

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/create" element={<CreateNote />} />
      <Route path="/notes/:id" element={<NoteDetail />} />
      <Route path="/notes/:id/edit" element={<EditNote />} />
      <Route path="*" element={<Navigate to="/" />} />
    </Routes>
  );
}