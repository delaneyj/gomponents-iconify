package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.84 22.73l-1.38-1.38c-.36.4-.88.65-1.46.65H6c-1.11 0-2-.89-2-2V8l1.06-1.05L1.11 3l1.28-1.27l19.72 19.73l-1.27 1.27M20 4a2 2 0 0 0-2-2h-8L7.6 4.4L20 16.8V4Z"/>`),
		g.Group(children),
	)
}