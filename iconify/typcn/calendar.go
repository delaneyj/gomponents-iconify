package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 6.184V6a3 3 0 1 0-6 0h-2a3 3 0 1 0-6 0v.184A2.997 2.997 0 0 0 3 9v9c0 1.654 1.346 3 3 3h12c1.654 0 3-1.346 3-3V9a2.997 2.997 0 0 0-2-2.816zM15 6a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0V6zM7 6a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0V6zm12 12c0 .551-.448 1-1 1H6c-.552 0-1-.449-1-1v-6h14v6z"/>`),
		g.Group(children),
	)
}