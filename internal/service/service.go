package service

import (
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/Vallghall/mt/test/internal/model/graph"
	"github.com/Vallghall/mt/test/internal/storage"
)

type GraphService interface {
	AddRelation(f *fact.Fact) error
	GetGraph() (graph.Graph, error)
}

type Service struct {
	GraphService
}

func NewService(s *storage.Storage) *Service {
	return &Service{
		GraphService: NewSocialService(s),
	}
}
