package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SensorWindowSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h16v20H4Zm3-11h3v-1h4v1h3V5H7v6Zm0 2v6h10v-6H7ZM6 4v16h12V4H6Z"/>`),
		g.Group(children),
	)
}