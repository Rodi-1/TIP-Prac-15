package service

type Service struct{ repo UserRepo }

func New(repo UserRepo) *Service { return &Service{repo: repo} }
func (s *Service) FindIDByEmail(email string) (int64, error) {
	u, e := s.repo.ByEmail(email)
	if e != nil {
		return 0, e
	}
	return u.ID, nil
}
