package store

var client Factory

type Factory interface {
	Employees() EmployeeStore
	Close() error
}

func Client() Factory {
	return client
}
func SetClient(factory Factory) {
	client = factory
}
