package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.5 16.25l2.86 1.69l-.75 1.22L15 17v-5h1.5v4.25M23 16c0 3.87-3.13 7-7 7c-2.38 0-4.47-1.19-5.74-3H4v-8H1l10-9l7 6.29c2.89.86 5 3.54 5 6.71m-2 0c0-2.76-2.24-5-5-5s-5 2.24-5 5s2.24 5 5 5s5-2.24 5-5Z"/>`),
		g.Group(children),
	)
}