package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ethereum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m5 1l-.09.31v8.88l.09.09l4-2.44L5 1Z"/><path fill="currentColor" d="M5 1L1 7.84l4 2.44V1Zm0 10.62l-.05.06v3.17L5 15l4-5.81l-4 2.44Z"/><path fill="currentColor" d="M5 15v-3.38L1 9.18l4 5.81Zm0-4.72l4-2.44l-4-1.87v4.31ZM1 7.84l4 2.44V5.97L1 7.84Z"/>`),
		g.Group(children),
	)
}