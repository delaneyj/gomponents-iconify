package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashAuto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.85 7.65L18 4l1.15 3.65M19 2h-2l-3.2 9h1.9l.7-2h3.2l.7 2h1.9M3 2v12h3v9l7-12H9l4-9H3Z"/>`),
		g.Group(children),
	)
}