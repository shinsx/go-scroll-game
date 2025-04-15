package audio

import (
	_ "embed"
)

var (
	//go:embed jab.wav
	Jab_wav []byte

	//go:embed jump.ogg
	Jump_ogg []byte
)
