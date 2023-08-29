package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.702 15.57c.457.383.901.559 2.01.559h.24a2.141 2.141 0 0 0 2.141-2.142h0a2.141 2.141 0 0 0-2.14-2.14h-2.251V9.5h4.139m-7.993.903l1.657-.903v6.629m-5.821 14.078c0 1.948 1.665 3.527 3.718 3.527h0c2.053 0 3.718-1.579 3.718-3.527v-2.292c0-1.948-1.665-3.527-3.718-3.527h0c-2.053 0-3.718 1.58-3.718 3.527m-.009-3.522V38.5m-3.913-9.481h-6.288m-1.567 4.674l4.717-14.056l4.716 14.098"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}