package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCombined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 10h-2V7a1 1 0 0 1 1-1a1 1 0 0 1 1 1zM9 6a1 1 0 0 0-1 1v1H7a1 1 0 0 0 0 2h3V7a1 1 0 0 0-1-1zm1 6v2H7a1 1 0 0 1-1-1a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M10 12v2H7a1 1 0 0 1-1-1a1 1 0 0 1 1-1Z"/><path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-3H7a1 1 0 0 1 0-2h3v-2H7a1 1 0 0 1 0-2h3v-2H7a1 1 0 0 1 0-2h1V7a1 1 0 0 1 2 0v3h2V7a1 1 0 0 1 2 0v3h2V7a1 1 0 0 1 2 0v3h3a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1Z" opacity=".5"/><path fill="currentColor" d="M10 16v2H7a1 1 0 0 1-1-1a1 1 0 0 1 1-1zm8-6h-2V7a1 1 0 0 1 1-1a1 1 0 0 1 1 1z"/>`),
		g.Group(children),
	)
}