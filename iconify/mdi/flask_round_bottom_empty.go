package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskRoundBottomEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15c0 3.87-3.13 7-7 7s-7-3.13-7-7c0-2.79 1.64-5.2 4-6.33V5c0-.55.45-1 1-1h.5l-1-2h5l-1 2h.5c.55 0 1 .45 1 1v3.67c2.36 1.13 4 3.54 4 6.33Z"/>`),
		g.Group(children),
	)
}