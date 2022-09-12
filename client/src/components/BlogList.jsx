import { Link } from 'react-router-dom'
import styled from 'styled-components'

import { useBlogs } from '../contexts/BlogsProvider'

const Container = styled.div`
  border: 2px #282828 solid;
  background-color: #1a1a1a;
  padding: 0.5em 2em;
`

const Item = styled.div`
  padding: 0.5em 1em;

  & + & {
    border-block-start: 1px solid #282828;
  }
`

const Title = styled.h2`
  font-size: 1.2em;
`

const StyledLink = styled(Link)`
  color: #2eadda;
  text-decoration: none;
`

const Blogs = () => {
  const { blogs } = useBlogs()

  return (
    <Container>
      {blogs.map((blog) => (
        <Item key={blog.id}>
          <Title>
            <StyledLink to={`/blogs/${blog.id}`}>
              {blog.title}
            </StyledLink>
          </Title>
        </Item>
      ))}
    </Container>
  )
}

export default Blogs
