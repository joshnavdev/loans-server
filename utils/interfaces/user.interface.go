package utils

type Repository[T comparable] interface {
  Create(model T) error
  FindOne(filter map[string]interface{}) (T, error)
  FindAll(filter map[string]interface{}) ([]T, error)
  FindById(id string) (T, error)
  Update(id string, model T) error
  Delete(id string) error
}
