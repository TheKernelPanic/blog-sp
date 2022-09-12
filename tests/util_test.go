package tests

import (
	"blog-sp-kernelpanic/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilenameGenerator(t *testing.T) {

	assert.Len(t, utils.FilenameGenerator("image/jpeg"), 36)
}
