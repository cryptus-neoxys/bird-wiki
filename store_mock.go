package main

import "github.com/stretchr/testify/mock"

// mock store contains additional methods for inspection
type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateBird(bird *Bird) error {
	/*
		When this method is called, `m.Called` records the call, and also
		returns the result that we pass to it (which you will see in the
		handler tests)
	*/
	rets := m.Called(bird)
	return rets.Error(0)
}

func (m *MockStore) GetBirds() ([] *Bird, error) {
	rets := m.Called()
	/*
		`rets.Get()` is a generic method, it returns whatever we pass to it,
		so we need to typecast it to the type we expect, that is -> []*Bird
	*/

	return rets.Get(0).([]*Bird), rets.Error(1)
}

func InitMockStore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}