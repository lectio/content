package content

import (
	"fmt"

	"github.com/stretchr/testify/suite"
)

const validFrontMatter = `
---
description: test description
---
test body
`

const noFrontMatter = `test body without front matter`

const invalidFrontMatter1 = `
---
description: test description

test body
`

type FrontMatterSuite struct {
	suite.Suite
}

func (suite *FrontMatterSuite) SetupSuite() {
}

func (suite *FrontMatterSuite) TearDownSuite() {
}

func (suite *FrontMatterSuite) TestNoFrontMatter() {
	fm := make(map[string]string)
	body, haveFrontMatter, err := ParseYAMLFrontMatter([]byte(noFrontMatter), &fm)
	suite.Nil(err, "Shouldn't have any errors")
	suite.False(haveFrontMatter, "Should not have any front matter")
	suite.Equal(fmt.Sprintf("%s", body), noFrontMatter)
}

func (suite *FrontMatterSuite) TestValidFrontMatter() {
	fm := make(map[string]string)
	body, haveFrontMatter, err := ParseYAMLFrontMatter([]byte(validFrontMatter), &fm)
	suite.Nil(err, "Shouldn't have any errors")
	suite.True(haveFrontMatter, "Should not front matter")

	suite.Equal(fmt.Sprintf("%s", body), "test body")

	descr, ok := fm["description"]
	suite.True(ok, "description should be found")
	suite.Equal(descr, "test description")
}

func (suite *FrontMatterSuite) TestInvalidFrontMatter() {
	fm := make(map[string]string)
	_, _, err := ParseYAMLFrontMatter([]byte(invalidFrontMatter1), &fm)
	suite.NotNil(err, "Should have error")
	suite.EqualError(err, "Unexplained front matter parser error; insideFrontMatter: true, yamlStartIndex: 5, yamlEndIndex: 0")
}
