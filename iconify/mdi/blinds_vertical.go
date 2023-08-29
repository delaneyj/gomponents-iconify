package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlindsVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 19V3H4v16H2v2h20v-2h-2m-10 0V5h4v14h-4Z"/>`),
		g.Group(children),
	)
}