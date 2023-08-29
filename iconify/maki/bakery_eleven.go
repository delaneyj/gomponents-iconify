package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BakeryEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4.5 2c-1 0-1 1-1 1L5 7.5s0 .5.5.5s.5-.5.5-.5L7.5 3s0-1-1-1h-2zM9 3.5l-2 4h1.5l1 1h.5c1 0 1-.9 1-.9V6.3L9 3.5zM0 6.3v1.2s.03 1.01 1 1c.97-.01.5 0 .5 0l1-1H4l-2-4l-2 2.8z" fill="currentColor"/>`),
		g.Group(children),
	)
}