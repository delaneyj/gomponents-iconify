package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10h-5V7c0-.6-.4-1-1-1H4V3c0-.6-.4-1-1-1c-.5 0-1 .4-1 1v18c0 .6.4 1 1 1s1-.4 1-1v-3h17c.6 0 1-.4 1-1v-6c0-.6-.4-1-1-1zm-7 0H4V8h10v2z"/>`),
		g.Group(children),
	)
}