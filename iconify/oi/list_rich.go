package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListRich(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v3h3V0H0zm4 0v1h4V0H4zm0 2v1h3V2H4zM0 4v3h3V4H0zm4 0v1h4V4H4zm0 2v1h3V6H4z"/>`),
		g.Group(children),
	)
}