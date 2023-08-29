package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CeilingLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9h3V4h2v5h3l4 8H4l4-8m6 9a2 2 0 0 1-2 2a2 2 0 0 1-2-2h4Z"/>`),
		g.Group(children),
	)
}