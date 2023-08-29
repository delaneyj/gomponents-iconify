package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="m6 25l5.171 15.628A2 2 0 0 0 13.07 42h21.86a2 2 0 0 0 1.899-1.372L42 25"/><path stroke-linecap="round" stroke-linejoin="round" d="M11.407 25.123h-5.09a.312.312 0 0 1-.313-.318C6.164 17.782 11.634 11.885 19 10h10c7.077 2.036 12.823 7.958 12.996 14.806a.31.31 0 0 1-.313.317h-5.09a9.56 9.56 0 0 0-6.297 2.366a9.56 9.56 0 0 1-12.592 0a9.56 9.56 0 0 0-6.297-2.366ZM19.1 10a5 5 0 1 1 9.8 0"/></g>`),
		g.Group(children),
	)
}