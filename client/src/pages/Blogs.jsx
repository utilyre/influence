import { useNavigate } from 'react-router-dom'
import styled from 'styled-components'

import BlogList from '../components/BlogList'
import { useUser } from '../contexts/UserProvider'
import Button from '../styled/Button'

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

const Home = () => {
  const navigate = useNavigate()

  const { user } = useUser()

  return (
    <div>
      <Heading>Blogs</Heading>

      <BlogsWrapper>
        <BlogList />

        {user !== null && (
          <Button onClick={() => navigate('/blogs/new')}>
            Create new blog
          </Button>
        )}
      </BlogsWrapper>
    </div>
  )
}

export default Home
