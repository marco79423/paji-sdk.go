package converter

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type converterTestSuite struct {
	suite.Suite
}

func (suite *converterTestSuite) TestConvertIntToIntRef() {
	var i = 10
	suite.Equal(&i, ConvertIntToIntRef(i))
}

func TestConverterTestSuite(t *testing.T) {
	suite.Run(t, new(converterTestSuite))
}
