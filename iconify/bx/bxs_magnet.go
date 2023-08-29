package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsMagnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M8 3H5a1 1 0 0 0-1 1v3h5V4a1 1 0 0 0-1-1zm7 1v3h5V4a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1zm0 10a3 3 0 0 1-6 0V9H4v5a8 8 0 0 0 16 0V9h-5v5z" fill="currentColor"/>`),
		g.Group(children),
	)
}