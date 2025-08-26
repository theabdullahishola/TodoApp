
import { useState } from "react";
import api from "../api/axios";
import { setAccessToken } from "../api/auth";
import { Link } from "react-router-dom";
import {  useGoogleLogin } from "@react-oauth/google";
import google from '../assets/google.png'

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");


 const googleLogin = useGoogleLogin({
  flow: "auth-code",  
  ux_mode: "popup",
  redirect_uri: import.meta.env.VITE_GOOGLE_REDIRECT_URI,
  onSuccess: async (resp) => {
    try {
   
      
      const res = await api.post("/api/auth/google", {code: resp.code});

      setAccessToken(res.data.accessToken); // your backend can still give you your own token
      window.location.href = "/todos";
    } catch (err) {
      setMessage(err.response?.data?.message || "Can't signup with Google");
    }
  },
  onError: () => setMessage("Google login failed"),
});



  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await api.post("/api/auth/login", { email, password });
      setAccessToken(res.data.accessToken);
      window.location.href = "/todos";
    } catch (err) {
      setMessage(err.response?.data?.message || "Login failed");
    }
  };

  return (
    <>
   
    <div>
      <h1 style={{textAlign: 'center'}}>Login</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <button type="submit">Login</button>
      </form>
      <div>
      <button onClick={googleLogin} className="btn-google">
      <img src={google} alt="Google" />
      </button>
      
    </div>
      <p style={{textAlign: 'center'}}> Don’t have an account?{" "}
  <Link to="/signup">Signup</Link></p>
      {message && <p style={{textAlign: 'center'}}>{message}</p>}
    </div>
    </>
   
  );
}
