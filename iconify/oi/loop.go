package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M6 0v1H1c-.55 0-1 .45-1 1v1h1V2h5v1l2-1.5L6 0zM2 4L0 5.5L2 7V6h5c.55 0 1-.45 1-1V4H7v1H2V4z"/>`),
		g.Group(children),
	)
}