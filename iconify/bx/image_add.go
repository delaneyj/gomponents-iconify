package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5h13v7h2V5c0-1.103-.897-2-2-2H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h8v-2H4V5z"/><path fill="currentColor" d="m8 11l-3 4h11l-4-6l-3 4z"/><path fill="currentColor" d="M19 14h-2v3h-3v2h3v3h2v-3h3v-2h-3z"/>`),
		g.Group(children),
	)
}