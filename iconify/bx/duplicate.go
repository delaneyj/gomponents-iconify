package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duplicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 10H9v3H6v2h3v3h2v-3h3v-2h-3z"/><path fill="currentColor" d="M4 22h12c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2zM4 8h12l.002 12H4V8z"/><path fill="currentColor" d="M20 2H8v2h12v12h2V4c0-1.103-.897-2-2-2z"/>`),
		g.Group(children),
	)
}