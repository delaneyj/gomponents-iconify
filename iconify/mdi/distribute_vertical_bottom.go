package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeVerticalBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3h10v5h5v2H2V8h5V3M2 19v2h20v-2h-3v-5H5v5H2Z"/>`),
		g.Group(children),
	)
}