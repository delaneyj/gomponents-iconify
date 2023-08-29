package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeHorizontalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2v3h2v14H8v3H6v-3H4V5h2V2h2m8 0v5h-2v10h2v5h2v-5h2V7h-2V2h-2Z"/>`),
		g.Group(children),
	)
}