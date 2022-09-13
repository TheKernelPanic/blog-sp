package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var charactersForRandom = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
var extensionMimetypeMap = map[string]string{
	"image/jpeg": "jpg",
	"image/png":  "png",
	"image/gif":  "gif",
}

func FilenameGenerator(mimetype string, seed int64) string {
	rand.Seed(seed * time.Now().UnixNano())
	b := make([]rune, 32)
	for i := range b {
		b[i] = charactersForRandom[rand.Intn(len(charactersForRandom))]
	}
	extension := extensionMimetypeMap[mimetype]
	return fmt.Sprintf("%s.%s", string(b), extension)
}
