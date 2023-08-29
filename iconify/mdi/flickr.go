package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M11 12c0 2.5-2 4.5-4.5 4.5S2 14.5 2 12s2-4.5 4.5-4.5S11 9.5 11 12m6.5-4.5C15 7.5 13 9.5 13 12s2 4.5 4.5 4.5s4.5-2 4.5-4.5s-2-4.5-4.5-4.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}