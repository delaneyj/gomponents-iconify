package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candelabra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 11c0 .55-.45 1-1 1H19v2c0 1.11-.89 2-2 2h-4v4h1a2 2 0 0 1 2 2H8c0-1.1.9-2 2-2h1v-4H7a2 2 0 0 1-2-2v-2h-.5c-.55 0-1-.45-1-1s.45-1 1-1H5V4l2 1v5h.5c.55 0 1 .45 1 1s-.45 1-1 1H7v2h4v-2h-.5c-.55 0-1-.45-1-1s.45-1 1-1h.5V2l2 1v7h.5c.55 0 1 .45 1 1s-.45 1-1 1H13v2h4v-2h-.5c-.55 0-1-.45-1-1s.45-1 1-1h.5V4l2 1v5h.5c.55 0 1 .45 1 1Z"/>`),
		g.Group(children),
	)
}