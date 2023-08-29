package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.897 4.042c.467-1.213-.725-2.405-1.938-1.938L2.962 10.179c-1.36.523-1.252 2.48.156 2.851l8.96 2.358a.75.75 0 0 1 .535.534l2.358 8.96c.37 1.408 2.328 1.516 2.85.157l8.076-20.997Zm-1.4-.538L16.421 24.5l-2.358-8.96a2.25 2.25 0 0 0-1.603-1.604L3.5 11.58l20.997-8.075Z"/>`),
		g.Group(children),
	)
}