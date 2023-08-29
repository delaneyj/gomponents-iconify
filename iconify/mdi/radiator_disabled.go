package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadiatorDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 12H4a2 2 0 0 0-2 2v8h2v-2h16v2h2v-8a2 2 0 0 0-2-2M7 17a1 1 0 0 1-1 1a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1a1 1 0 0 1 1 1v2m4 0a1 1 0 0 1-1 1a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1a1 1 0 0 1 1 1v2m4 0a1 1 0 0 1-1 1a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1a1 1 0 0 1 1 1v2m4 0a1 1 0 0 1-1 1a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1a1 1 0 0 1 1 1v2Z"/>`),
		g.Group(children),
	)
}