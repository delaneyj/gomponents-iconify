package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductSubscriptions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 8H4V6h16Zm-2-6H6v2h12Zm4 10v8a2 2 0 0 1-2 2H4a2.006 2.006 0 0 1-2-2v-8a2.006 2.006 0 0 1 2-2h16a2.006 2.006 0 0 1 2 2Zm-8.073 5.042l2.323-1.986l-3.059-.256L12 12l-1.191 2.8l-3.059.256l2.323 1.986l-.7 2.958L12 18.428L14.627 20Z"/>`),
		g.Group(children),
	)
}