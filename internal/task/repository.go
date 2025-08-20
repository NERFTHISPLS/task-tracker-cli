package task

type Repository interface {
	Add(task Task) error
	Update(task Task) error
	Delete(id int) error
	List() ([]Task, error)
	ListByStatus(status string) ([]Task, error)
	ByID(id int) (Task, error)
}
