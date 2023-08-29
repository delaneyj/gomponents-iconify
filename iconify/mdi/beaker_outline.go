package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeakerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3h18v2a2 2 0 0 0-2 2v12a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2V3m4 2v2h5v1H7v1h3v1H7v1h3v1H7v1h5v1H7v1h3v1H7v3h10V5H7Z"/>`),
		g.Group(children),
	)
}