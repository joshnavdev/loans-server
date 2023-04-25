package common

func GetValue[K comparable] (value, defaultValue K) K {
  if isNil(value) {
    return defaultValue
  }

  return value
}

func isNil[T comparable] (value T) bool {
  var t T

  return value == t
}
