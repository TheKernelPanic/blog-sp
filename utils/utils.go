package utils

import (
	"fmt"
	"math/rand"
)

var charactersForRandom = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
var extensionMimetypeMap = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/gif":  "gif",
}

func FilenameGenerator(mimetype string) string {
	b := make([]rune, 32)
	for i := range b {
		b[i] = charactersForRandom[rand.Intn(len(charactersForRandom))]
	}
	extension := extensionMimetypeMap[mimetype]
	return fmt.Sprintf("%s.%s", string(b), extension)
}
