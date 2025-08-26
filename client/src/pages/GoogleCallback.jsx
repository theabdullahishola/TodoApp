// src/pages/GoogleCallback.jsx
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import api from "../api/axios";
import { setAccessToken } from "../api/auth";

export default function GoogleCallback() {
  const navigate = useNavigate();

  useEffect(() => {
    const urlParams = new URLSearchParams(window.location.search);
    const code = urlParams.get("code");

    if (code) {
      (async () => {
        try {
          const res = await api.post("/api/auth/google", { code });
          setAccessToken(res.data.accessToken);
          navigate("/todos");
        } catch (err) {
          console.error("Google login failed:", err);
          navigate("/login");
        }
      })();
    } else {
      // No code? Just go back to login
      navigate("/login");
    }
  }, [navigate]);

  return <p>Logging in with Google...</p>;
}
