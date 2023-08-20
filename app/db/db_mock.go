package db

// MockDB is a mock implementation of the DB interface.
type MockDB struct {
	PingFunc  func() error
	CloseFunc func() error
}

func (m *MockDB) Ping() error {
	return m.PingFunc()
}

func (m *MockDB) Close() error {
	return m.CloseFunc()
}
