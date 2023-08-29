package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerThreeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19v-3h6v-2.5H9v-3h5V8H8V5h9v5.5L15.5 12l1.5 1.5V19H8Z"/>`),
		g.Group(children),
	)
}