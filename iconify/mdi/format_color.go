package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M9.624 11.998L12 5.665l2.375 6.333m-3.375-9l-5.5 14h2.25l1.125-3h6.25l1.125 3h2.25l-5.5-14h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}