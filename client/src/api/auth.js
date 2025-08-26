import api from "./axios"

export const setAccessToken = (token) => {
  localStorage.setItem("access_token", token);
};

export const getAccessToken = () => {
  return localStorage.getItem("access_token");
};

export const logout = async (setIsLoggingOut) => {
  setIsLoggingOut(true)
  try {
    // ask backend to clear refresh cookie
    await api.post("/api/auth/logout");
  } catch (err) {
    console.error("Logout error:", err);
  } finally {
    // clear stored access token
    localStorage.removeItem("access_token");
    
    window.location.href = "/login";
    setIsLoggingOut(false)
  }
};
