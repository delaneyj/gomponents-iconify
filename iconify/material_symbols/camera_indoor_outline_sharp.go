package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraIndoorOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17h6v-2l2 1.05v-4.1L14 13v-2H8v6Zm-4 4V9l8-6l8 6v12H4Zm2-2h12v-9l-6-4.5L6 10v9Zm6-6.75Z"/>`),
		g.Group(children),
	)
}