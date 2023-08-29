package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Icecream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 17h4V8A7 7 0 1 0 5 8v9h4v3a3 3 0 1 0 6 0v-3Zm2-2V8A5 5 0 0 0 7 8v7h4v5a1 1 0 1 0 2 0v-5h4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}