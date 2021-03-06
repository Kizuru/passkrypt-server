package registering

type Service interface {
	RegisterUser(...User) error
}

type Repository interface {
	RegisterUser(User) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) RegisterUser(users ...User) error {
	for _, user := range users {
		if err := s.r.RegisterUser(user); err != nil {
			return err
		}
	}

	return nil
}
