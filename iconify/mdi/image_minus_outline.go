package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.4 14.21c-.93.7-1.65 1.67-2.05 2.79H6.5l2.75-3.53l1.96 2.36l2.75-3.54l1.44 1.92M5 19V5h14v8c.7 0 1.37.13 2 .35V5a2 2 0 0 0-2-2H5c-1.1 0-2 .9-2 2v14a2 2 0 0 0 2 2h8.35c-.22-.63-.35-1.3-.35-2H5m10-1v2h8v-2h-8Z"/>`),
		g.Group(children),
	)
}