import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import Auth from '../components/auth/auth.tsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <Auth isAuth={false}/>
  </StrictMode>,
)
