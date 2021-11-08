package imginfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEtag(t *testing.T) {
	assert := assert.New(t)
	etag, err := GetEtag("../testdata/cLena.bmp")
	assert.Nil(err)
	assert.Equal(etag, "FkNeoBle1hbDNe71uzruZSc3Qv-W")
}
