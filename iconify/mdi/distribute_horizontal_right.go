package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17V7h5V2h2v20H8v-5H3m16 5h2V2h-2v3h-5v14h5v3Z"/>`),
		g.Group(children),
	)
}