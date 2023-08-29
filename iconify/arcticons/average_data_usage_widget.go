package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AverageDataUsageWidget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.986 29.62a9 9 0 0 1-10.56 2.6a9 9 0 0 1-5.354-9.464a9 9 0 0 1 7.667-7.712m2.678.026a9 9 0 0 1 6.425 4.458a9 9 0 0 1 .512 7.803"/>`),
		g.Group(children),
	)
}