package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pushpin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 4H16l4 3l-4 13s-6 4-6 8h10l4 16l4-16h10c0-4-4-6.833-6-8L28 7l4-3Z"/>`),
		g.Group(children),
	)
}