package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.9 2.3l18.804 6.27a.5.5 0 0 1 .028.938L13 13l-4.425 8.85a.5.5 0 0 1-.928-.086L2.261 2.912a.5.5 0 0 1 .638-.612Z"/>`),
		g.Group(children),
	)
}