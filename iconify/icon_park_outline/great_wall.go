package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreatWall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 9v31h40V9h-8v7h-8V9h-8v7h-8V9H4Zm0 15h40M4 32h40m-20-8v8m-8 0v8m0-24v8m16 8v8m0-24v8"/>`),
		g.Group(children),
	)
}