package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCombined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2H3c-.6 0-1 .4-1 1v18c0 .6.4 1 1 1h6c.6 0 1-.4 1-1v-3H7c-.6 0-1-.4-1-1s.4-1 1-1h3v-2H7c-.6 0-1-.4-1-1s.4-1 1-1h3v-2H7c-.6 0-1-.4-1-1s.4-1 1-1h1V7c0-.6.4-1 1-1s1 .4 1 1v3h2V7c0-.6.4-1 1-1s1 .4 1 1v3h2V7c0-.6.4-1 1-1s1 .4 1 1v3h3c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}