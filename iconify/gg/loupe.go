package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loupe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M11 11V8h2v3h3v2h-3v3h-2v-3H8v-2h3Z"/><path fill-rule="evenodd" d="M3 12a9 9 0 0 0 9 9h6a3 3 0 0 0 3-3v-6a9 9 0 1 0-18 0Zm9-7a7 7 0 1 1 0 14a7 7 0 0 1 0-14Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}