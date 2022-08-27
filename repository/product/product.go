package product

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repo *Repository) GetData() {
}
