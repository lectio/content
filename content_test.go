package content

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(FrontMatterSuite))
}
