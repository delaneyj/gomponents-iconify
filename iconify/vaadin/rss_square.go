package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm3.6 14c-.9 0-1.6-.7-1.6-1.6s.7-1.6 1.6-1.6s1.6.7 1.6 1.6S4.6 14 3.6 14zm4 0c0-3.1-2.5-5.6-5.6-5.6V6c4.4 0 8 3.6 8 8H7.6zm4 0c0-5.3-4.3-9.6-9.6-9.6V2c6.6 0 12 5.4 12 12h-2.4z"/>`),
		g.Group(children),
	)
}