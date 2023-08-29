package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxlProductHunt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M13.337 9h-2.838v3h2.838a1.501 1.501 0 1 0 0-3z" fill="currentColor"/><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm1.337 12h-2.838v3H8.501V7h4.837a3.498 3.498 0 0 1 3.499 3.499a3.499 3.499 0 0 1-3.5 3.501z" fill="currentColor"/>`),
		g.Group(children),
	)
}