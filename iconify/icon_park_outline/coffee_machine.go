package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 42h29a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H17"/><path d="M22 36c5.523 0 10-4.477 10-10H12c0 5.523 4.477 10 10 10Z"/><path stroke-linecap="round" d="M16 18v2m6-2v2m6-2v2"/></g>`),
		g.Group(children),
	)
}