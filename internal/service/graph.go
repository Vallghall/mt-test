package service

import (
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/Vallghall/mt/test/internal/model/graph"
	"github.com/Vallghall/mt/test/internal/storage"
)

type SocialService struct {
	s *storage.Storage
}

func NewSocialService(s *storage.Storage) *SocialService {
	return &SocialService{s}
}

func (s *SocialService) AddRelation(f *fact.Fact) error {
	return s.s.StoreRelation(f)
}

func (s *SocialService) GetGraph() (graph.Graph, error) {
	return s.s.GetGraph()
}
