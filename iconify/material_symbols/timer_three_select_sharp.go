package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerThreeSelectSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-3h6v-2.5H4v-3h6V8H4V5h9v5.5L11.5 12l1.5 1.5V19H4Zm11 0v-2h4v-1h-4v-5h6v2h-4v1h4v5h-6Z"/>`),
		g.Group(children),
	)
}