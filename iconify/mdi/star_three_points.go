package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarThreePoints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2.6l-3 9.8l-7 7.5l10-2.3L22 20l-7-7.5l-3-9.9Z"/>`),
		g.Group(children),
	)
}