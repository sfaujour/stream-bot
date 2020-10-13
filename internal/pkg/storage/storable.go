package storage

type storable interface {
	put(key string, value string) error
	get(key string) (string, error)
}
