package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraIndoorSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V9l8-6l8 6v12H4Zm4-4h6v-2l2 1.05v-4.1L14 13v-2H8v6Z"/>`),
		g.Group(children),
	)
}