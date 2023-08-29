package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopologyFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 18a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM8 18a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM8 6a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm12 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM6 8v8m12 0V8M8 6h8m0 12H8M7.5 7.5l9 9m-9 0l9-9"/>`),
		g.Group(children),
	)
}