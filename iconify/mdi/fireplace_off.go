package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireplaceOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 22H2v-2h20v2m0-16H2V3h20v3m-2 1v12h-3v-8s-2.5-1-5-1s-5 1-5 1v8H4V7h16Z"/>`),
		g.Group(children),
	)
}