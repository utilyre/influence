import { useContext, createContext } from 'react'
import PropTypes from 'prop-types'

import useLocalStorage from '../hooks/useLocalStorage'

const UserContext = createContext()

export const useUser = () => {
  const ctx = useContext(UserContext)
  return ctx
}

const UserProvider = ({ children }) => {
  const [user, setUser] = useLocalStorage('user', null)

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}
    </UserContext.Provider>
  )
}

UserProvider.propTypes = {
  children: PropTypes.element,
}

export default UserProvider
