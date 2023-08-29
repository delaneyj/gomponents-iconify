package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 8l8-5l4 2.5M4 8v8l8 5M4 8l4 2.5m4 2.5l8-5m-8 5v8m0-8l-4-2.5M20 8v8l-8 5m8-13l-4-2.5m-8 5l8-5"/>`),
		g.Group(children),
	)
}