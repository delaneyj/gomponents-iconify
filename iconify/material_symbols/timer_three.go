package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19v-3h6v-2.5H9v-3h5V8H8V5h6q1.25 0 2.125.875T17 8v1.9q0 .875-.613 1.488T14.9 12q.875 0 1.488.613T17 14.1V16q0 1.25-.875 2.125T14 19H8Z"/>`),
		g.Group(children),
	)
}