package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func One(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2Zm-8.049 29.085H38.5m-6.049-21.17H38.5M32.451 24h3.944m-3.944-10.585v21.17"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.928 34.585v-21.17l8.015 21.17v-21.17M9.5 30.578a4.007 4.007 0 1 0 8.015 0V17.422a4.007 4.007 0 1 0-8.015 0Zm2.007-12.489l2-1.089m0 0v14"/>`),
		g.Group(children),
	)
}