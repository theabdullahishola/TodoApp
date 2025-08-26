import { useState } from "react";
import api from "../api/axios";
import { Link } from "react-router-dom";
import { useGoogleLogin } from "@react-oauth/google";
import google from '../assets/google.png'
import { setAccessToken } from "../api/auth";

export default function Signup() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");


  const googleLogin = useGoogleLogin({
    flow: "auth-code",   // or "auth-code" if you want to exchange later
    ux_mode: "popup",
    redirect_uri: import.meta.env.VITE_GOOGLE_REDIRECT_URI,
    onSuccess: async (resp) => {
      try {
        // resp contains id_token now
        
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
      const res = await api.post("/api/auth/signup", { email, password });
      setMessage(res.data.message);
      window.location.href = "/login";
    } catch (err) {
      setMessage(err.response?.data?.message || "Error signing up");// why this
    }
  };

  return (
    <>
    <div>
      <h1 style={{textAlign: 'center'}}>Signup</h1>
      
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
                <button type="submit" className="btn-signup">Signup</button>
            </form>

    </div>

            <div>
                <button onClick={googleLogin} className="btn-google">
                <img src={google} alt="Google" />
                </button>
            </div>
  
    
    
    <p style={{textAlign: 'center'}}>Have an account?{" "}
  <Link to="/login">Login</Link></p>

      {message && <p style={{textAlign: 'center'}}>{message}</p>}
   
    </>
  );
}
