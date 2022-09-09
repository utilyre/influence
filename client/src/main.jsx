import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'

import './main.css'
import App from './App'
import BlogsProvider from './contexts/BlogsProvider'
import UserProvider from './contexts/UserProvider'

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <BrowserRouter>
      <BlogsProvider>
        <UserProvider>
          <App />
        </UserProvider>
      </BlogsProvider>
    </BrowserRouter>
  </React.StrictMode>
)
