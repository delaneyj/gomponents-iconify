package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsDiamondOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2L5 12l7 10l7-10M7.44 12L12 5.5l4.56 6.5L12 18.5"/>`),
		g.Group(children),
	)
}