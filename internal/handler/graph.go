package handler

import (
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addFact(c *gin.Context) {
	var f fact.Fact
	err := c.BindJSON(&f)
	catch(c, err, http.StatusInternalServerError)

	if f.A == f.B {
		catch(c, ErrSamePerson, http.StatusBadRequest)
	}

	if f.A.IsValid() && f.B.IsValid() && f.Communication.DescriptionNoNEmpty() {
		catch(c, ErrBadRequest, http.StatusBadRequest)
	}

	err = h.s.AddRelation(&f)
}

func (h *Handler) getRelationsGraph(c *gin.Context) {

}
