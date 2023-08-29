package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h14v1h6v14h-8v-1H4v6H2V2Zm2 12h10V4H4v10Zm12-9v10h4V5h-4Z"/>`),
		g.Group(children),
	)
}