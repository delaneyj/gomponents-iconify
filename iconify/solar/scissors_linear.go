package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M16.401 20.5L6 2m16 17a3 3 0 1 1-6 0a3 3 0 0 1 6 0ZM7.599 20.5L18 2M2 19a3 3 0 1 0 6 0a3 3 0 0 0-6 0Z"/>`),
		g.Group(children),
	)
}