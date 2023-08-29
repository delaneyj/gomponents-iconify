package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOscillating(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14h3l-4 4l-4-4h3c0-2.7 1.7-7.4 7-7.9v2c-3.4.5-5 3.8-5 5.9m14 0c0-2.7-1.7-7.4-7-7.9v2c3.4.6 5 3.8 5 5.9h-3l4 4l4-4h-3Z"/>`),
		g.Group(children),
	)
}