package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.9 17.2c-.6 0-1-.4-1-1c0-.3.1-.5.3-.7l3.5-3.5l-3.5-3.5c-.4-.4-.4-1 0-1.4c.4-.4 1-.4 1.4 0l4.2 4.2c.4.4.4 1 0 1.4l-4.2 4.2c-.2.2-.5.3-.7.3z"/>`),
		g.Group(children),
	)
}