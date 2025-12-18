package service

import (
    "context"
    "go-user-service/db/sqlc"
    "go-user-service/internal/repository"
    "time"
)

type UserService struct {
    repo *repository.Repo
}

func NewUserService(r *repository.Repo) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) CalculateAge(dob time.Time) int {
    now := time.Now()
    age := now.Year() - dob.Year()
    if now.YearDay() < dob.YearDay() {
        age--
    }
    return age
}

// Helper to access Repo directly if needed
func (s *UserService) Repo() *repository.Repo { 
    return s.repo 
}

func (s *UserService) Create(ctx context.Context, name string, dob string) (int64, error) {
    dobTime, _ := time.Parse("2006-01-02", dob)
    res, err := s.repo.Q.CreateUser(ctx, db.CreateUserParams{Name: name, Dob: dobTime})
    if err != nil { return 0, err }
    return res.LastInsertId()
}