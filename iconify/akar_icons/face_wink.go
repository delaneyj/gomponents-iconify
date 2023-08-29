package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceWink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M7 9h2m7 .05v-.1"/><path stroke-linejoin="round" d="M16 15c-.5 1-1.79 2-4 2s-3.5-1-4-2"/></g>`),
		g.Group(children),
	)
}