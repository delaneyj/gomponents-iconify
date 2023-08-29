package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeathStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.05 13h19.9c-.5 5.05-4.76 9-9.95 9c-5.18 0-9.45-3.95-9.95-9m19.9-2H2.05c.5-5.05 4.77-9 9.95-9s9.45 3.95 9.95 9M12 6.75a2.5 2.5 0 0 0-2.5-2.5A2.5 2.5 0 0 0 7 6.75a2.5 2.5 0 0 0 2.5 2.5a2.5 2.5 0 0 0 2.5-2.5Z"/>`),
		g.Group(children),
	)
}