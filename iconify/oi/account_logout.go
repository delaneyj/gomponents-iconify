package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountLogout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3 0v1h4v5H3v1h5V0H3zM2 2L0 3.5L2 5V4h4V3H2V2z"/>`),
		g.Group(children),
	)
}