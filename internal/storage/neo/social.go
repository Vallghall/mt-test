package neo

import (
	"encoding/json"
	"fmt"
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/Vallghall/mt/test/internal/model/graph"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"os"
)

type SocialStorage struct {
	s neo4j.Session
}

func NewSocialStorage(s neo4j.Session) *SocialStorage {
	return &SocialStorage{s}
}

func (s *SocialStorage) StoreRelation(f *fact.Fact) error {
	query :=
		`MERGE 
			(a:PERSON {name: $name1, surname: $surname1, age: $age1}) 
			-[rel:SENDS_MESSAGE_TO {info: $info}]-> 
			(b:PERSON {name: $name2, surname: $surname2, born: $age2})`
	_, err := s.s.Run(query, map[string]any{
		"name1":    f.A.Name,
		"surname1": f.A.Surname,
		"age1":     f.A.Age,
		"info":     f.Communication.Description,
		"name2":    f.B.Name,
		"surname2": f.B.Surname,
		"age2":     f.B.Age,
	})
	return err
}

func (s *SocialStorage) LoadGraphData(src string) error {
	f, err := os.Open(src)
	if err != nil {
		log.Fatalln(err)
	}
	var g graph.Graph
	err = json.NewDecoder(f).Decode(&g)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fact := range g.Facts {
		if err = s.StoreRelation(&fact); err != nil {
			log.Println(err)
		}
	}
	return nil
}

func (s *SocialStorage) GetGraph() (graph.Graph, error) {
	query := `match (n:PERSON) -[rel:SENDS_MESSAGE_TO]-> (m:PERSON) return n, rel, m`
	result, err := s.s.Run(query, map[string]any{})
	if err != nil {
		return graph.Graph{}, err
	}

	_ = make([]fact.Fact, 0)
	for result.Next() {
		//var f fact.Fact
		fmt.Println(result.Record().Values)
	}
	return graph.Graph{}, nil
}
