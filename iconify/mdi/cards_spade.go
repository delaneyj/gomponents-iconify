package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsSpade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C9 7 4 9 4 14c0 2 2 4 4 4c1 0 2 0 3-1c0 0 .32 2-2 5h6c-2-3-2-5-2-5c1 1 2 1 3 1c2 0 4-2 4-4c0-5-5-7-8-12Z"/>`),
		g.Group(children),
	)
}