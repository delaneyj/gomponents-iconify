package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Puzzle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M4 24V12h9v-2a6 6 0 0 1 12 0v2h9v12h4a6 6 0 0 1 0 12h-4v8H4v-8h4a6 6 0 0 0 0-12H4Z"/>`),
		g.Group(children),
	)
}