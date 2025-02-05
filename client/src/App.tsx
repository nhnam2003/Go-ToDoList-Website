import axios from "axios";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Login from "./Page/LoginPage";
import RegisterPage from "./Page/Register";
import TodoPage from "./Page/TodoPage/TodoPage";
import Layout from "./components/Layout";

export const BASE_URL =
  import.meta.env.MODE === "development" ? "http://localhost:5000/api" : "/api";
axios.defaults.baseURL = "http://localhost:5000/api";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route element={<Layout></Layout>}>
            <Route path="/login" index element={<Login></Login>}></Route>
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/dashboard" element={<TodoPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
