package repository

type DataManager interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Remove(key string) error
}
