package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sandstorm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 22h32a8 8 0 1 0-8-8M10 29H4m23 0h-6m23 0h-6m-28 6H4m23 0h-6m23 0h-6m-22 7H4m40 0H32"/>`),
		g.Group(children),
	)
}