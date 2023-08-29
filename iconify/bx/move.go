package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 11h-5V6h3l-4-4l-4 4h3v5H6V8l-4 4l4 4v-3h5v5H8l4 4l4-4h-3v-5h5v3l4-4l-4-4z"/>`),
		g.Group(children),
	)
}