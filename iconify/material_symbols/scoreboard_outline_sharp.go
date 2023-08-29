package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScoreboardOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 15V9H19v6h-4.5Zm1.5-1.5h1.5v-3H16v3ZM5 15v-3.5h3v-1H5V9h4.5v3.5h-3v1h3V15H5Zm6.25-4V9.5h1.5V11h-1.5Zm0 3.5V13h1.5v1.5h-1.5ZM2 20V4h5V2h2v2h6V2h2v2h5v16H2Zm2-2h7.25v-1.5h1.5V18H20V6h-7.25v1.5h-1.5V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}