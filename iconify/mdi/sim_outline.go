package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SimOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4v16H6V8.8L10.8 4H18m0-2h-8L4 8v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2M9.5 19h-2v-2h2v2m7 0h-2v-2h2v2m-7-4h-2v-4h2v4m3.5 4h-2v-4h2v4m0-6h-2v-2h2v2m3.5 2h-2v-4h2v4Z"/>`),
		g.Group(children),
	)
}