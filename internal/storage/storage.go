package storage

import (
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/Vallghall/mt/test/internal/model/graph"
	"github.com/Vallghall/mt/test/internal/storage/neo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type GraphStorage interface {
	StoreRelation(fact *fact.Fact) error
	LoadGraphData(src string) error
	GetGraph() (graph.Graph, error)
}

type Storage struct {
	GraphStorage
}

func NewStorage(s neo4j.Session) *Storage {
	return &Storage{
		GraphStorage: neo.NewSocialStorage(s),
	}
}
