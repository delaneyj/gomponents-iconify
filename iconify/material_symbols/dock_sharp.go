package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 23v-2h8v2H8Zm-2-4V1h12v18H6Zm2-5h8V6H8v8Z"/>`),
		g.Group(children),
	)
}