package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h2v2H3zm0 4h2v2H3zm0 4h2v2H3zM3 3h2v2H3zm8 0h2v2h-2zM7 3h2v2H7zm8 0h2v2h-2zm4 0h2v2h-2zm0 12h2v2h-2zm0-4h2v2h-2zm0-4h2v2h-2zm-4 4h2v2h-2zm-8 0h2v2H7zm4-4h2v2h-2zm0 8h2v2h-2zm0-4h2v2h-2zm6 8H3v2h18v-2h-2z"/>`),
		g.Group(children),
	)
}