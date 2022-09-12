import { createContext, useContext, useState } from 'react'
import PropTypes from 'prop-types'

const BlogsContext = createContext()

export const useBlogs = () => {
  const ctx = useContext(BlogsContext)
  return ctx
}

const BlogsProvider = ({ children }) => {
  const [lastId, setLastId] = useState(10)
  const [blogs, setBlogs] = useState([
    {
      id: 1,
      title: 'count all objects within a values_list Django',
      content:
        'Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis.',
    },

    {
      id: 2,
      title: 'Can I create view from two columns with the same foreign key?',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },

    {
      id: 3,
      title: 'matplotlib put legend side-by-side',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },

    {
      id: 4,
      title:
        'Apple Developer and Distribution Certification Expiring - Proper Way to Regenerate',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },

    {
      id: 5,
      title:
        'Clickable ListTile in Flutter to take to location based on JSON coordinates',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },

    {
      id: 6,
      title: 'Powershell - Sort files into directory based on file name',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },

    {
      id: 7,
      title: 'How to edit Google Maps entry for a town?',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },

    {
      id: 8,
      title:
        'Why Makefile is compiling a file type FILE while should be a .exe',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },

    {
      id: 9,
      title: 'Fully disable WooCommerce Endpoints',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },

    {
      id: 10,
      title:
        'Are there alternatives to create a merchant profile on Google Play when your country is not listed?',
      content:
        'Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.',
    },
  ])

  const createBlog = (title, content) => {
    setBlogs((prev) => [...prev, { id: lastId + 1, title, content }])
    setLastId((prev) => prev + 1)
  }

  return (
    <BlogsContext.Provider value={{ blogs, createBlog }}>
      {children}
    </BlogsContext.Provider>
  )
}

BlogsProvider.propTypes = {
  children: PropTypes.any,
}

export default BlogsProvider
