package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortBoolDescending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 7h3l-4-4l-4 4h3v14h2M5 7c0 1.1.9 2 2 2s2-.9 2-2s-.9-2-2-2s-2 .9-2 2m2-4c2.21 0 4 1.79 4 4s-1.79 4-4 4s-4-1.79-4-4s1.79-4 4-4m0 10c-2.21 0-4 1.79-4 4s1.79 4 4 4s4-1.79 4-4s-1.79-4-4-4Z"/>`),
		g.Group(children),
	)
}