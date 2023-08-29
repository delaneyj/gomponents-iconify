package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceHappy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 9.05v-.1m8 .1v-.1"/><path stroke-linejoin="round" d="M16 14c-.5 1.5-1.79 3-4 3s-3.5-1.5-4-3"/></g>`),
		g.Group(children),
	)
}