import { useParams } from 'react-router-dom'

import { useBlogs } from '../contexts/BlogsProvider'
import toKebabcase from '../utils/toKebabcase'

const Blog = () => {
  const { title } = useParams()

  const { blogs } = useBlogs()

  const blog = blogs.find((blog) => toKebabcase(blog.title) === title)

  return (
    <div>
      <h1>{blog.title}</h1>
      <p>{blog.content}</p>
    </div>
  )
}

export default Blog
