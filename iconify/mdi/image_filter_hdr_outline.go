package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageFilterHdrOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 18H1l7.25-9.67l2 2.67L14 6l9 12m-11.5-5.33L14 16h5l-5-6.67l-2.5 3.34M5 16h6.5l-3.25-4.33L5 16Z"/>`),
		g.Group(children),
	)
}