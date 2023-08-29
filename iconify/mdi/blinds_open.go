package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlindsOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h18c.55 0 1 .45 1 1v2c0 .55-.45 1-1 1h-1v1c0 .55-.45 1-1 1h-6v2.17A3 3 0 0 1 15 13c0 1.66-1.34 3-3 3s-3-1.34-3-3c0-1.31.84-2.42 2-2.83V8H5c-.55 0-1-.45-1-1V6H3c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1m9 10c-.55 0-1 .45-1 1s.45 1 1 1s1-.45 1-1s-.45-1-1-1Z"/>`),
		g.Group(children),
	)
}