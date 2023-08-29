package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 12h-3a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1zm7-7H3a1 1 0 1 0 0 2h17a1 1 0 1 0 0-2zm-2 3H5a1 1 0 0 0-1 1v8c0 1.654 1.346 3 3 3h9c1.654 0 3-1.346 3-3V9a1 1 0 0 0-1-1zm-2 10H7c-.552 0-1-.449-1-1v-7h11v7c0 .551-.448 1-1 1z"/>`),
		g.Group(children),
	)
}