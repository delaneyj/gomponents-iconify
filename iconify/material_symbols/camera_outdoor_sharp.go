package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOutdoorSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V9l8-6l8 6v2h-9v8h9v2H4Zm8-3v-6h6v2l2-1.05v4.1L18 16v2h-6Z"/>`),
		g.Group(children),
	)
}