package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 23v-2h8v2H8Zm-2-4V1h12v18H6Zm2-3v1h8v-1H8Zm0-2h8V6H8v8ZM8 4h8V3H8v1Zm0 0V3v1Zm0 12v1v-1Z"/>`),
		g.Group(children),
	)
}