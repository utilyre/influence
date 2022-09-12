import { NavLink } from 'react-router-dom'
import styled from 'styled-components'
import { useUser } from '../contexts/UserProvider'

const HorizontalList = styled.ul`
  display: flex;
  list-style: none;
  justify-content: space-around;
  align-items: center;
  background-color: #161616;
  margin: 0;
  height: 4em;

  border-bottom: #1b1b1b solid 2px;
`

const Item = styled.li`
  display: flex;
  gap: 1em;
`

const StyledNavLink = styled(NavLink)`
  text-decoration: none;
  color: #ffffff;
  font-weight: bold;
  font-size: 1.1rem;
  transition: 300ms ease color;

  &.active {
    color: #2eadda;
  }
`

const Navigation = () => {
  const { user } = useUser()

  return (
    <HorizontalList>
      <Item>
        <StyledNavLink to='/'>Home</StyledNavLink>
      </Item>

      <Item>
        <StyledNavLink to='/blogs'>Blogs</StyledNavLink>
      </Item>

      {user === null && (
        <Item>
          <StyledNavLink to='/signin'>Sign In</StyledNavLink>
          <StyledNavLink to='/signup'>Sign Up</StyledNavLink>
        </Item>
      )}
    </HorizontalList>
  )
}

export default Navigation
