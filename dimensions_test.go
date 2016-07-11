package imginfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDimension(t *testing.T) {
	assert := assert.New(t)

	w, h, err := Dimension("./testdata/lena.jpg")
	assert.Nil(err)
	assert.Equal(w, 512)
	assert.Equal(h, 512)

	w, h, err = Dimension("./testdata/lena.png")
	assert.Nil(err)
	assert.Equal(w, 512)
	assert.Equal(h, 512)
}
