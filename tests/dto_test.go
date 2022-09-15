package tests

import (
	"blog-sp-kernelpanic/dto"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshalJSONWithHtmlSection(t *testing.T) {

	var serialized string

	postDescription := "Foo"
	postSlug := "foo"
	sectionContent := "<p>Foo</p>"

	serialized = "{\"description\":\"Foo\",\"slug\":\"foo\",\"sections\":[{\"mimetype\":\"text/html\",\"content\":\"<p>Foo</p>\",\"sort\":1}]}"

	var postDto dto.Post

	err := json.Unmarshal([]byte(serialized), &postDto)
	if err != nil {
		t.Fail()
	}

	assert.Len(t, postDto.Sections, 1)
	assert.Equal(t, postDescription, postDto.Description)
	assert.Equal(t, postSlug, postDto.Slug)

	for _, sectionDto := range postDto.Sections {
		switch sectionType := sectionDto.(type) {
		case *dto.TextSection:
			assert.Equal(t, sectionContent, sectionType.Content)
			assert.Equal(t, 1, sectionType.Sort)
		default:
			t.Fail()
		}
	}
}
