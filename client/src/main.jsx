import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { GoogleOAuthProvider } from "@react-oauth/google";

const queryClient = new QueryClient();

createRoot(document.getElementById('root')).render(
  
  <StrictMode>
     <GoogleOAuthProvider clientId={import.meta.env.VITE_GOOGLE_CLIENT_ID}>
     <QueryClientProvider client={queryClient}>
     <App />
     </QueryClientProvider>

     </GoogleOAuthProvider>

    
  </StrictMode>,
)
