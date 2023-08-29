package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsRegistered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M13 9h-3v2h3a1 1 0 0 0 0-2z" fill="currentColor"/><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2zm2.13 15l-2.67-4H10v4H8V7h5a3 3 0 0 1 .79 5.88L16.54 17z" fill="currentColor"/>`),
		g.Group(children),
	)
}