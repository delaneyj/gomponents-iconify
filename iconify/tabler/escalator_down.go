package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.5 7h2.733a2 2 0 0 1 1.337.513L18 16h1.5a2.5 2.5 0 1 1 0 5h-2.733a2 2 0 0 1-1.337-.513L6 12H4.5a2.5 2.5 0 1 1 0-5zM18 3v7m-3-3l3 3l3-3"/>`),
		g.Group(children),
	)
}