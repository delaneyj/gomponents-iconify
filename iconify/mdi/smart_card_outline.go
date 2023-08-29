package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartCardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 3H2a2.07 2.07 0 0 0-2 2v14a2.07 2.07 0 0 0 2 2h20a2.07 2.07 0 0 0 2-2V5a2.07 2.07 0 0 0-2-2m0 16H2V5h20m-8 12v-1.25c0-1.66-3.34-2.5-5-2.5s-5 .84-5 2.5V17h10M9 7a2.5 2.5 0 1 0 2.5 2.5A2.5 2.5 0 0 0 9 7m6 3v3h4v-3h-4"/>`),
		g.Group(children),
	)
}