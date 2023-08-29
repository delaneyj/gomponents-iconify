package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 1h16v22H4V1Zm2 2v18h12V3H6Zm5 14h2.004v2.004H11V17Z"/>`),
		g.Group(children),
	)
}