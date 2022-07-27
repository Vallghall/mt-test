package neo

import (
	"encoding/json"
	"fmt"
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/Vallghall/mt/test/internal/model/graph"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
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
	query := fmt.Sprintf(
		`MERGE
			(a:PERSON {name: $name1, surname: $surname1, age: $age1})
		 MERGE
			(b:PERSON {name: $name2, surname: $surname2, age: $age2})
		 MERGE (a)-[:%s {info: $info}]->(b)`, f.Communication.Type)
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
	query :=
		`MATCH (n:PERSON) -[rel:SENDS_MESSAGE_TO]-> (m:PERSON) 
		RETURN n AS first, rel AS relation, m AS second`
	result, err := s.s.Run(query, map[string]any{})
	if err != nil {
		return graph.Graph{}, err
	}

	g := make([]fact.Fact, 0)
	for result.Next() {
		vals := result.Record()
		personA, _ := vals.Get("first")
		personB, _ := vals.Get("second")
		relation, _ := vals.Get("relation")
		a, b, rel := personA.(dbtype.Node), personB.(dbtype.Node), relation.(dbtype.Relationship)
		f := fact.Fact{
			A: fact.Person{
				Name:    a.Props["name"].(string),
				Surname: a.Props["surname"].(string),
				Age:     int(a.Props["age"].(int64)),
			},
			B: fact.Person{
				Name:    b.Props["name"].(string),
				Surname: b.Props["surname"].(string),
				Age:     int(b.Props["age"].(int64)),
			},
			Communication: fact.Communication{
				Type:        rel.Type,
				Description: rel.Props["info"].(string),
			},
		}
		g = append(g, f)
	}
	return graph.Graph{g}, nil
}
