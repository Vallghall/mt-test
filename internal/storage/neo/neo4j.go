package neo

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

type Config struct {
	URI      string
	Username string
	Password string
}

func NewNeo4jSession(c *Config) (neo4j.Session, func()) {
	driver, err := neo4j.NewDriver(c.URI, neo4j.BasicAuth(c.Username, c.Password, ""))
	if err != nil {
		log.Fatalln(err)
	}

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	return session, func() {
		driver.Close()
	}
}
