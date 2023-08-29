package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDEffects(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M11 3h2l9 3v11l-10 3.5L2 17V6.5L4 6l9 2.5l-2 .5l-9-2.5"/>`),
		g.Group(children),
	)
}