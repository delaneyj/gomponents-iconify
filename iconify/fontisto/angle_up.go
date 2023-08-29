package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.57 19.2L24 16.778L12 4.8L0 16.778L2.43 19.2L12 9.653z"/>`),
		g.Group(children),
	)
}