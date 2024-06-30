package service

import (
	//	"goproject_SBG-bot/repository"
	//	"fmt"
	"goproject_SBG-bot/datastruct"
)

type repo interface {
	Chek_avtorisation(chatID int64) bool
	Out_list(chatID int64) string
	Get(chatID string) datastruct.Person
	Get_previous(chatID int64) string
	Get_worker() ([][]int64, []string)
	Сancel(chatID int64) string
	AddName(chatID int64) error
	AddNameWork(text_vvod string, chatID int64) error
	DeleteName(chatID int64) (string, error)
	DeleteNameWork(text_vvod string, chatID int64) error
	EnterName(text_vvod string, chatID int64) error
	EnterDate(text_vvod string, chatID int64) error
}

type Service struct {
	Repo repo
}

func New(r repo) *Service {

	return &Service{
		Repo: r,
	}
}

func (s *Service) Chek_avtorisation(chatID int64) bool {
	return s.Repo.Chek_avtorisation(chatID)
}

func (s *Service) Out_list(chatID int64) string {
	return s.Repo.Out_list(chatID)
}

func (s *Service) AddName(chatID int64) error {
	return s.Repo.AddName(chatID)
}

func (s *Service) AddNameWork(text_vvod string, chatID int64) error {
	return s.Repo.AddNameWork(text_vvod, chatID)
}

func (s *Service) DeleteName(chatID int64) (string, error) {
	return s.Repo.DeleteName(chatID)
}

func (s *Service) DeleteNameWork(text_vvod string, chatID int64) error {
	return s.Repo.DeleteNameWork(text_vvod, chatID)
}

func (s *Service) Сancel(chatID int64) string {
	return s.Repo.Сancel(chatID)
}

func (s *Service) Get(chatID string) datastruct.Person {
	return s.Repo.Get(chatID)
}

func (s *Service) Get_worker() ([][]int64, []string) {
	return s.Repo.Get_worker()
}
func (s *Service) Get_previous(chatID int64) string {
	return s.Repo.Get_previous(chatID)
}

func (s *Service) EnterName(text_vvod string, chatID int64) error {
	return s.Repo.EnterName(text_vvod, chatID)
}

func (s *Service) EnterDate(text_vvod string, chatID int64) error {
	return s.Repo.EnterDate(text_vvod, chatID)
}
