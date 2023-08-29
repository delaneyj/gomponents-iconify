package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsChevronUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm4.293 12.707L12 10.414l-4.293 4.293l-1.414-1.414L12 7.586l5.707 5.707l-1.414 1.414z" fill="currentColor"/>`),
		g.Group(children),
	)
}