// src/api/api.js
import axios from "axios";

const BASE_URL = import.meta.env.MODE==="development" ? "http://localhost:8080" : ""; // your backend

// Axios instance
const api = axios.create({
  baseURL: BASE_URL,
  withCredentials: true, // important: send cookies
});


// Request interceptor to attach access token

api.interceptors.request.use(
  async (config) => {
    const token = localStorage.getItem("access_token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor to handle expired access token
// Response interceptor
api.interceptors.response.use(
    (response) => response,
    async (error) => {
      const originalRequest = error.config;
  
      // If unauthorized AND not already retried AND not the refresh endpoint
      if (
        error.response &&
        error.response.status === 401 &&
        !originalRequest._retry &&
        !originalRequest.url.includes("/api/auth/refresh")
      ) {
        originalRequest._retry = true;
        try {
          // call refresh using api (cookies attached)
          const res = await api.get("/api/auth/refresh");
          localStorage.setItem("access_token", res.data.accessToken);
  
          // update header and retry
          originalRequest.headers["Authorization"] = `Bearer ${res.data.accessToken}`;
          return api(originalRequest);
        } catch (err) {
          localStorage.removeItem("access_token");
          window.location.href = "/login";
          return Promise.reject(err);
        }
      }
  
      return Promise.reject(error);
    }
  );

  export default api;
  