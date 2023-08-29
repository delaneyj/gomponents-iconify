package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M248 48v208h48V58.7c23.9 13.8 40 39.7 40 69.3v128h48V128C384 57.3 326.7 0 256 0h-64C121.3 0 64 57.3 64 128v128h48V128c0-29.6 16.1-55.5 40-69.3V256h48V48h48zM48 288c-12.1 0-23.2 6.8-28.6 17.7l-16 32c-5 9.9-4.4 21.7 1.4 31.1S20.9 384 32 384v96c0 17.7 14.3 32 32 32s32-14.3 32-32v-96h256v96c0 17.7 14.3 32 32 32s32-14.3 32-32v-96c11.1 0 21.4-5.7 27.2-15.2s6.4-21.2 1.4-31.1l-16-32c-5.4-10.9-16.5-17.7-28.6-17.7H48z"/>`),
		g.Group(children),
	)
}