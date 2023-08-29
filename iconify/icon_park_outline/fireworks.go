package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fireworks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="m6 42l8.674-24.736L31 34.038L6 42Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m23 19l5-5c2.667-2.667 3-5 1-7m0 18l5-5c3.333-3.333 6.667-3.333 10 0"/><path fill="currentColor" d="M20 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm22-1a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 23a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-3 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g>`),
		g.Group(children),
	)
}