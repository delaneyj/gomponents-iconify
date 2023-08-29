package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 5h8.42l1.16 2H19v9h-6l-1.15-2H6v7H5V5m13 10V8h-4l-1.15-2H6v7h6.42l1.16 2H18Z"/>`),
		g.Group(children),
	)
}