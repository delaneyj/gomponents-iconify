package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SensorWindowOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v20H4Zm2-11h4v-1h4v1h4V4H6v7Zm0 2v7h12v-7H6Zm0 7h12H6Z"/>`),
		g.Group(children),
	)
}