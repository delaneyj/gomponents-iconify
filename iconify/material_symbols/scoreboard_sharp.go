package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScoreboardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 15H19V9h-4.5v6Zm1.5-1.5v-3h1.5v3H16ZM5 15h4.5v-1.5h-3v-1h3V9H5v1.5h3v1H5V15Zm6.25-4h1.5V9.5h-1.5V11Zm0 3.5h1.5V13h-1.5v1.5ZM2 20V4h5V2h2v2h6V2h2v2h5v16H2Zm9.25-12.5h1.5V6h-1.5v1.5Zm0 10.5h1.5v-1.5h-1.5V18Z"/>`),
		g.Group(children),
	)
}