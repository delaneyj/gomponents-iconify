package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardFortyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 3c2.6 0 5 1 6.9 2.6L21 3v7h-7l2.6-2.6c-1.4-1.2-3.2-1.9-5.1-1.9C8 5.5 4.9 7.8 3.9 11l-2.4-.8C2.9 6 6.8 3 11.5 3m1.5 9h6v2h-4v2h2c1.1 0 2 .9 2 2v2c0 1.1-.9 2-2 2h-4v-2h4v-2h-4v-6m-8 0v6h4v4h2V12H9v4H7v-4H5Z"/>`),
		g.Group(children),
	)
}