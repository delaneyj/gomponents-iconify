package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallerText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.75 18h-1.51a.64.64 0 0 1-.42-.13a.83.83 0 0 1-.24-.32l-1-2.65H7.41l-1 2.65a.79.79 0 0 1-.23.31a.62.62 0 0 1-.42.14H4.25L9 6h2zm-3.69-4.5L10.4 9.12a12.13 12.13 0 0 1-.4-1.3q-.09.39-.2.72t-.2.58L7.95 13.5z"/>`),
		g.Group(children),
	)
}