package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M41 4H7a3 3 0 0 0-3 3v34a3 3 0 0 0 3 3h34a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="M12 19.054c3.325-4 7.325-6 12-6s8.675 2 12 6"/><path fill="currentColor" d="M24 31a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="4" d="m19 21l5.008 7"/></g>`),
		g.Group(children),
	)
}