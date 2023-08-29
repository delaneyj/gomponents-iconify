package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 0L1 2v1h9V2L5.5 0zM2 4v4L1 9v1h9V9L9 8V4H2zm1.49 1a.5.5 0 0 1 .36.15L5.5 6.79l1.65-1.64A.5.5 0 0 1 8 5.5v3a.5.5 0 0 1-1 0V6.71L5.85 7.85a.5.5 0 0 1-.707.003L5.14 7.85L4 6.71V8.5a.5.5 0 0 1-1 0v-3a.5.5 0 0 1 .49-.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}