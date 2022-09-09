import { Routes, Route } from 'react-router-dom'
import styled from 'styled-components'

import Navigation from './components/Navigation'
import Login from './pages/Login'
import Home from './pages/Home'
import Blogs from './pages/Blogs'
import Blog from './pages/Blog'
import NewBlog from './pages/NewBlog'

const Nav = styled.nav`
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
`

const Main = styled.main`
  margin-block-start: 6em;
`

const App = () => {
  return (
    <>
      <Nav>
        <Navigation />
      </Nav>

      <Main>
        <Routes>
          <Route path='/' element={<Home />} />
          <Route path='/login' element={<Login />} />
          <Route path='/blogs' element={<Blogs />} />
          <Route path='/blogs/:title' element={<Blog />} />
          <Route path='/blogs/new' element={<NewBlog />} />
        </Routes>
      </Main>
    </>
  )
}

export default App
