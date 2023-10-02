package intelonemono

import (
	"sync"

	"gio.tools/fonts/intelonemono/intelonemonobold"
	"gio.tools/fonts/intelonemono/intelonemonobolditalic"
	"gio.tools/fonts/intelonemono/intelonemonoitalic"
	"gio.tools/fonts/intelonemono/intelonemonolight"
	"gio.tools/fonts/intelonemono/intelonemonolightitalic"
	"gio.tools/fonts/intelonemono/intelonemonomedium"
	"gio.tools/fonts/intelonemono/intelonemonomediumitalic"
	"gio.tools/fonts/intelonemono/intelonemonoregular"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var (
	once       sync.Once
	collection []font.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(intelonemonobold.OTF)
		register(intelonemonobolditalic.OTF)
		register(intelonemonoitalic.OTF)
		register(intelonemonolight.OTF)
		register(intelonemonolightitalic.OTF)
		register(intelonemonomedium.OTF)
		register(intelonemonomediumitalic.OTF)
		register(intelonemonoregular.OTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(src []byte) {
	faces, err := opentype.ParseCollection(src)
	if err != nil {
		panic("failed to parse font: " + err.Error())
	}
	collection = append(collection, faces[0])
}
