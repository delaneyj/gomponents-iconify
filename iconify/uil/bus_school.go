package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusSchool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 12.5v4a1 1 0 0 0 1 1h1a3 3 0 0 0 6 0h6a3 3 0 0 0 6 0h1a1 1 0 0 0 1-1v-10a3 3 0 0 0-3-3H8.44A3 3 0 0 0 5.6 5.55L4.16 9.86l-2.71 1.81a1 1 0 0 0-.45.83Zm20-3h-2v-4h1a1 1 0 0 1 1 1Zm-4 8a1 1 0 1 1 1 1a1 1 0 0 1-1-1Zm-2-6h6v4h-.78a3 3 0 0 0-4.44 0H15Zm0-6h2v4h-2Zm-4 6h2v4h-2Zm0-6h2v4h-2Zm-2 4H6.39l1.1-3.32a1 1 0 0 1 .95-.68H9Zm-4 8a1 1 0 1 1 1 1a1 1 0 0 1-1-1ZM3 13l2.3-1.5H9v4h-.78a3 3 0 0 0-4.44 0H3Z"/>`),
		g.Group(children),
	)
}