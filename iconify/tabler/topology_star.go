package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopologyStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 18a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM20 6a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM8 6a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm12 12a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-6-6a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM7.5 7.5l3 3m-3 6l3-3m3 0l3 3m0-9l-3 3"/>`),
		g.Group(children),
	)
}