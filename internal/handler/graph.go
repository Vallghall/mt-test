package handler

import (
	"github.com/Vallghall/mt/test/internal/model/fact"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func (h *Handler) addFact(c *gin.Context) {
	var f fact.Fact
	err := c.BindJSON(&f)
	catch(c, err, http.StatusInternalServerError)

	if ok, err := regexp.MatchString(``, regexp.QuoteMeta(f.Communication.Type)); !ok || err != nil {
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

func (h *Handler) getRelationsGraph(c *gin.Context) {
	g, err := h.s.GetGraph()
	catch(c, err, http.StatusInternalServerError)

	c.JSON(http.StatusOK, g)
}
