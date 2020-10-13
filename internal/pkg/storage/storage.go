package storage

type Storage struct {
	Engine storable
}

func New(engine storable) *Storage {
	return &Storage{
		Engine: engine,
	}
}

func (s *Storage) Put(key string, value string) error {
	return s.Engine.put(key, value)
}

func (s *Storage) Get(key string) (string, error) {
	return s.Engine.get(key)
}
