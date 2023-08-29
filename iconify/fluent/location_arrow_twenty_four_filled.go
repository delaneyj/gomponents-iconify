package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.896 4.042c.467-1.213-.725-2.405-1.938-1.938L2.962 8.64c-1.36.522-1.252 2.48.156 2.85l7.011 1.845a.75.75 0 0 1 .535.535l1.845 7.01c.37 1.409 2.328 1.516 2.85.157l6.537-16.996Z"/>`),
		g.Group(children),
	)
}