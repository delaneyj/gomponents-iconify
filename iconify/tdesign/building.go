package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h10v3h3v4h2v2h-1v9h1v2H2v-2h1v-9H2V9h2V5h3V2Zm2 3h6V4H9v1Zm-4 6v9h3v-6h8v6h3v-9H5Zm13-2V7H6v2h12Zm-4 11v-4h-4v4h4Z"/>`),
		g.Group(children),
	)
}