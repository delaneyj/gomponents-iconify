package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircuitChangeover(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h2m16-5h2M4 12a2 2 0 1 0 4 0a2 2 0 1 0-4 0m12-5a2 2 0 1 0 4 0a2 2 0 1 0-4 0m4 10h2m-6 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-8.5-6.5L16 7"/>`),
		g.Group(children),
	)
}