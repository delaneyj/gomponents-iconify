package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.43 4.8L0 7.222L12 19.2L24 7.222L21.57 4.8L12 14.347z"/>`),
		g.Group(children),
	)
}