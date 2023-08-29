package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpForum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 6h-3v9H6v3h12l4 4V6zm-5 7V2H2v15l4-4h11z"/>`),
		g.Group(children),
	)
}