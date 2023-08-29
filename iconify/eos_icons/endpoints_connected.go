package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndpointsConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.02 13.01v-2h-4.01v-5h1v-4h-4v4h1v5h-12v2H7V18H6v4h3.99v-4h-1v-4.99h12.03z"/>`),
		g.Group(children),
	)
}