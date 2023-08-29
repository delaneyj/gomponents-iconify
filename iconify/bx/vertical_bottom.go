package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 19h3v2h-3zm-5 0h3v2h-3zm-5 0h3v2H8zm-5 0h3v2H3zM13 5h-2v8H8l4 4l4-4h-3V5z"/>`),
		g.Group(children),
	)
}