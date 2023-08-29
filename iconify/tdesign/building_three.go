package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2h8v3h3v4h3v13H2V9h3V5h3V2Zm2 3h4V4h-4v1ZM7 9h10V7H7v2Zm-3 2v9h4v-6h8v6h4v-9H4Zm10 9v-4h-4v4h4Z"/>`),
		g.Group(children),
	)
}