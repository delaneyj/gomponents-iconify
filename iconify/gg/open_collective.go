package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenCollective(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".5" d="m16.682 15.753l2.13 2.13A8.965 8.965 0 0 0 21 12a8.964 8.964 0 0 0-2.123-5.806l-2.133 2.133A5.974 5.974 0 0 1 18 12c0 1.42-.493 2.725-1.318 3.753Z"/><path d="M15.673 16.744a6 6 0 1 1 .08-9.426l2.13-2.13a9 9 0 1 0-.077 13.689l-2.133-2.133Z"/></g>`),
		g.Group(children),
	)
}