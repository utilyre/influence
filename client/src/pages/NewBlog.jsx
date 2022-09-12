import { useEffect, useRef, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import styled from 'styled-components'

import toKebabcase from '../utils/toKebabcase'
import { useBlogs } from '../contexts/BlogsProvider'
import Button from '../styled/Button'

const Form = styled.form`
  display: flex;
  flex-direction: column;
  padding: 2em;
`

const Label = styled.label`
  margin-block-end: 0.25em;
`

const Input = styled.input`
  height: 2em;
  font-size: 1rem;
  font-weight: bold;
  padding-inline: 0.5em;
  border: 1px solid #282828;
  background-color: #1a1a1a;
  margin-block-end: 1em;
  outline: none;

  &:focus {
    border-color: #2f2f2f;
  }
`

const Textarea = styled.textarea`
  min-height: 20em;
  padding: 0.5em;
  border: 1px solid #282828;
  background-color: #1a1a1a;
  margin-block-end: 1em;
  outline: none;
  resize: vertical;

  &:focus {
    border-color: #2f2f2f;
  }
`

const NewBlog = () => {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  const titleRef = useRef()
  const navigate = useNavigate()

  const { createBlog } = useBlogs()

  const submitHandler = (e) => {
    e.preventDefault()

    createBlog(title, content)
    navigate(`/blogs/${toKebabcase(title)}`, {
      replace: true,
    })
  }

  useEffect(() => {
    titleRef.current?.focus()
  }, [])

  return (
    <Form onSubmit={submitHandler}>
      <Label>Title</Label>
      <Input
        type='text'
        required
        ref={titleRef}
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />

      <Label>Content</Label>
      <Textarea
        required
        value={content}
        onChange={(e) => setContent(e.target.value)}
      ></Textarea>

      <Button type='submit' className='centered'>
        Create
      </Button>
    </Form>
  )
}

export default NewBlog
