package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SizeXl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 7h2l1 2.5L10 7h2l-2 5l2 5h-2l-1-2.5L8 17H6l2-5l-2-5m7 0h2v8h4v2h-6V7Z"/>`),
		g.Group(children),
	)
}