package graph

import "github.com/Vallghall/mt/test/internal/model/fact"

// Graph is a social graph where vertices are people
// and the edges are relations between them
type Graph struct {
	Facts []fact.Fact `json:"facts"`
}
