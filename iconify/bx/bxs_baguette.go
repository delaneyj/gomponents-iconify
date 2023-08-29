package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsBaguette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M11.13 4.41l4.23 4.23L14.3 9.7l-4.24-4.24l-1.77 1.77l4.24 4.24l-1.06 1.06l-4.24-4.24l-1.77 1.77L9.7 14.3l-1.06 1.06l-4.23-4.23C1.86 14 1.55 18 3.79 20.21a5.38 5.38 0 0 0 3.85 1.5a8 8 0 0 0 5.6-2.47l6-6c2.87-2.87 3.31-7.11 1-9.45s-6.24-1.93-9.11.62z" fill="currentColor"/>`),
		g.Group(children),
	)
}