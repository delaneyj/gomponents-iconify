package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyVisualizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h3v2H3Zm0-4v-2h8v2H3Zm0-4v-2h18v2H3Zm0-4V7h8v2H3Zm0-4V3h3v2H3Zm5 16v-2h3v2H8ZM8 5V3h3v2H8Zm5 16v-2h3v2h-3Zm0-4v-2h8v2h-8Zm0-8V7h8v2h-8Zm0-4V3h3v2h-3Zm5 16v-2h3v2h-3Zm0-16V3h3v2h-3Z"/>`),
		g.Group(children),
	)
}