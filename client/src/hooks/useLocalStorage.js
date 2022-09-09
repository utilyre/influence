import { useState } from 'react'

const PREFIX = 'influence-'

const useLocalStorage = (key, init) => {
  const prefixedKey = PREFIX + key

  const [storedValue, setStoredValue] = useState(() => {
    const item = localStorage.getItem(prefixedKey)
    return item ? JSON.parse(item) : init
  })

  const setValue = (value) => {
    const valueToStore = value instanceof Function ? value(storedValue) : value
    setStoredValue(valueToStore)
    localStorage.setItem(prefixedKey, JSON.stringify(valueToStore))
  }

  return [storedValue, setValue]
}

export default useLocalStorage
