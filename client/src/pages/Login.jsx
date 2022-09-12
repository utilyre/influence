import { useEffect, useRef, useState } from 'react'
import PropTypes from 'prop-types'
import { Navigate } from 'react-router-dom'
import styled from 'styled-components'

import { useUser } from '../contexts/UserProvider'

const Form = styled.form`
  display: flex;
  flex-direction: column;
  max-width: 300px;
  margin-inline: auto;

  height: 100vh;

  align-items: center;
  justify-content: center;
`

const Label = styled.label`
  align-self: start;
`

const Input = styled.input`
  margin-block-end: 2em;
  height: 2em;
  background: none;
  border: none;
  border-bottom: 2px solid #6c6c6c;
  outline: none;
  transition: 300ms ease border-bottom-color;
  width: 100%;

  &:hover {
  }

  &:focus {
    border-bottom-color: #d95024;
  }
`

const Button = styled.button`
  background-color: #3e83e3;
  border: none;
  border-radius: 0.25em;
  height: 2.5em;
  width: 10em;
  margin-inline: auto;
  cursor: pointer;
  transition: 300ms ease background-color;
  margin-block-start: 1em;

  &:hover {
    background-color: #5391e6;
  }

  &:focus {
    outline: #3e83e3 solid 3px;
    outline-offset: 0.5em;
  }
`

const Login = ({ signUp: isSignUp }) => {
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const nameRef = useRef()
  const emailRef = useRef()

  const { user, setUser } = useUser()

  const submitHandler = (e) => {
    e.preventDefault()

    setUser({
      id: 1, // TODO: avoid using hardcoded id
      name: name,
      email: email,
      password: password,
    })
  }

  useEffect(() => {
    if (isSignUp) {
      nameRef.current?.focus()
    } else {
      emailRef.current?.focus()
    }
  }, [])

  return user === null ? (
    <Form onSubmit={submitHandler}>
      {isSignUp && (
        <>
          <Label>Name</Label>
          <Input
            type='text'
            required
            ref={nameRef}
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </>
      )}

      <Label>Email</Label>
      <Input
        type='email'
        required
        ref={emailRef}
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />

      <Label>Password</Label>
      <Input
        type='password'
        required
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      <Button type='submit'>Login</Button>
    </Form>
  ) : (
    <Navigate replace to='/' />
  )
}

Login.propTypes = {
  signUp: PropTypes.bool,
}

export default Login
