import { useMemo } from 'react'
import { useParams } from 'react-router-dom'

import { useBlogs } from '../contexts/BlogsProvider'

const Blog = () => {
  const { id } = useParams()
  const { blogs } = useBlogs()

  const blog = useMemo(() => blogs.find((b) => b.id == id), [id, blogs])

  return (
    <div>
      <h1>{blog?.title}</h1>
      <p>{blog?.content}</p>
    </div>
  )
}

export default Blog
