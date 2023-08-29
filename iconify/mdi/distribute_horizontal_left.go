package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7v10h-5v5h-2V2h2v5h5M5 2H3v20h2v-3h5V5H5V2Z"/>`),
		g.Group(children),
	)
}