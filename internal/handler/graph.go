package handler

import (
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/Vallghall/mt/test/internal/model/graph"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func (h *Handler) addFact(c *gin.Context) {
	var f fact.Fact
	err := c.BindJSON(&f)
	catch(c, err, http.StatusInternalServerError)

	if ok, err := regexp.MatchString(`^([A-Z]+-?[A-Z]+)+$`, regexp.QuoteMeta(f.Communication.Type)); !ok || err != nil {
		catch(c, ErrInvalidCommType, http.StatusBadRequest)
	}

	if f.A == f.B {
		catch(c, ErrSamePerson, http.StatusBadRequest)
	}

	if !(f.A.IsValid() && f.B.IsValid() && f.Communication.DescriptionNoNEmpty()) {
		catch(c, ErrBadRequest, http.StatusBadRequest)
	}

	err = h.s.AddRelation(&f)
	catch(c, err, http.StatusInternalServerError)
	c.JSON(http.StatusOK, map[string]string{
		"message": "OK",
	})
}

type PersonWithRelationsCount struct {
	fact.Person
	Count int `json:"count"`
}

type Response struct {
	PersonWithMaxRelations PersonWithRelationsCount `json:"person_with_max_relations"`
	PersonWithMinRelations PersonWithRelationsCount `json:"person_with_min_relations"`
	AverageRelationsAmount int                      `json:"average_relations_amount"`
	graph.Graph
}

func (h *Handler) getRelationsGraph(c *gin.Context) {
	g, err := h.s.GetGraph()
	catch(c, err, http.StatusInternalServerError)

	m := make(map[fact.Person]int)
	var r Response
	for _, val := range g.Facts {
		if _, ok := m[val.A]; !ok {
			m[val.A] = 0
		}
		m[val.A]++
	}

	min, max, count := 999, 0, 0
	for k, v := range m {
		count++
		r.AverageRelationsAmount += v
		if v > max {
			max = v
			r.PersonWithMaxRelations.Person = k
			r.PersonWithMaxRelations.Count = v
		}

		if v < min {
			min = v
			r.PersonWithMinRelations.Person = k
			r.PersonWithMinRelations.Count = v
		}
	}

	r.AverageRelationsAmount /= count
	r.Graph = g

	c.JSON(http.StatusOK, r)
}
