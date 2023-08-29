package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsChevronLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm2.707 14.293l-1.414 1.414L7.586 12l5.707-5.707l1.414 1.414L10.414 12l4.293 4.293z" fill="currentColor"/>`),
		g.Group(children),
	)
}