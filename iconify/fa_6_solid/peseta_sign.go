package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PesetaSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 32c-17.7 0-32 14.3-32 32v96c-17.7 0-32 14.3-32 32s14.3 32 32 32v224c0 17.7 14.3 32 32 32s32-14.3 32-32v-96h96c77.4 0 142-55 156.8-128h3.2c17.7 0 32-14.3 32-32s-14.3-32-32-32h-3.2C334 87 269.4 32 192 32H64zm218.5 128H96V96h96c41.8 0 77.4 26.7 90.5 64zM96 224h186.5c-13.2 37.3-48.7 64-90.5 64H96v-64z"/>`),
		g.Group(children),
	)
}