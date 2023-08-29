package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M0 0v4h4V0H0zm5 2v3H2v1h4V2H5zm2 2v3H4v1h4V4H7z"/>`),
		g.Group(children),
	)
}