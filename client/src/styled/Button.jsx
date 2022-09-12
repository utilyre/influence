import styled from 'styled-components'

const Button = styled.button`
  background-color: #3e83e3;
  border: none;
  border-radius: 0.25em;
  height: 2.5em;
  width: 10em;
  cursor: pointer;
  transition: 300ms ease background-color;
  margin-block-start: 1em;

  &.centered {
    margin-inline: auto;
  }

  &:hover {
    background-color: #5391e6;
  }

  &:focus {
    outline: #3e83e3 solid 3px;
    outline-offset: 0.5em;
  }
`

export default Button
