package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneDoorFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 19h10V5H7v14zm6-8h2v2h-2v-2z" opacity=".3"/><path fill="currentColor" d="M13 11h2v2h-2z"/><path fill="currentColor" d="M19 19V5c0-1.1-.9-2-2-2H7c-1.1 0-2 .9-2 2v14H3v2h18v-2h-2zm-2 0H7V5h10v14z"/>`),
		g.Group(children),
	)
}