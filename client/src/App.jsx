import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import { useEffect, useState,createContext } from "react";
import api from "./api/axios";
import { getAccessToken, setAccessToken } from "./api/auth";
import GoogleCallback from "./pages/GoogleCallback";

// Pages
import Login from "./pages/Login";
import Signup from "./pages/SignUp";
import TodoApp from "./pages/TodoDashboard";

export const AuthContext = createContext();
export default function App() {
  const [loading, setLoading] = useState(true);    
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoggingOut, setIsLoggingOut] = useState(false);
  

  async function checkAuth() {
    if (isLoggingOut) {
      setLoading(false);
      return;
    }

    const token = getAccessToken();
    if (token) {
      setIsAuthenticated(true);
      setLoading(false);
      return;
    }
    try {
      const res = await api.get("/api/auth/refresh");
      setAccessToken(res.data.accessToken);
      setIsAuthenticated(true);
    } catch {
      setIsAuthenticated(false);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    checkAuth();
  }, [isLoggingOut]);



  return (
    <AuthContext.Provider value={{ isLoggingOut, setIsLoggingOut }}>
    <Router>
      <Routes>
        {/* Public routes */}
        <Route
          path="/login"
          element={isAuthenticated ? <Navigate to="/todos" replace /> : <Login />}
        />
        <Route
          path="/signup"
          element={isAuthenticated ? <Navigate to="/todos" replace /> : <Signup />}
        />
        <Route path="/auth/callback" element={<GoogleCallback />} />
        {/* Private route */}
        <Route
          path="/todos"
          element={isAuthenticated ? <TodoApp /> : <Navigate to="/login" replace />}
        />

        {/* Default route */}
        <Route
          path="*"
          element={<Navigate to={isAuthenticated ? "/todos" : "/login"} replace />}
        />
      </Routes>
    </Router>
    </AuthContext.Provider>
  );
}
