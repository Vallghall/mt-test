package service

import (
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/Vallghall/mt/test/internal/storage"
)

type SocialService struct {
	s *storage.Storage
}

func NewSocialService(s *storage.Storage) *SocialService {
	return &SocialService{s}
}

func (s *SocialService) AddRelation(f *fact.Fact) error {
	//TODO implement me
	panic("implement me")
}

func (s *SocialService) GetGraph() {
	//TODO implement me
	panic("implement me")
}
