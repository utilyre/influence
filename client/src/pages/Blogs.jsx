import { useNavigate } from 'react-router-dom'
import styled from 'styled-components'

import BlogList from '../components/BlogList'
import { useUser } from '../contexts/UserProvider'

const Heading = styled.h1`
  text-align: center;
`

const BlogsWrapper = styled.div`
  max-width: 850px;
  margin-inline: auto;
  margin-block: 2em;
  margin-inline-min: 2em;

  display: flex;
  flex-direction: column;
`

const NewBlog = styled.button`
  align-self: flex-end;
  margin-block-start: 2em;
  margin-inline-end: 1em;
  height: 3em;
  padding-inline: 1em;
  background-color: #8e32ec;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: 300ms ease background-color;

  &:hover {
    background-color: #9a48ee;
  }

  &:focus {
    outline: 3px #8e32ec solid;
    outline-offset: 0.5em;
  }
`

const Home = () => {
  const navigate = useNavigate()

  const { user } = useUser()

  return (
    <div>
      <Heading>Blogs</Heading>

      <BlogsWrapper>
        <BlogList />

        {user !== null && (
          <NewBlog onClick={() => navigate('/blogs/new')}>
            Create new blog
          </NewBlog>
        )}
      </BlogsWrapper>
    </div>
  )
}

export default Home
