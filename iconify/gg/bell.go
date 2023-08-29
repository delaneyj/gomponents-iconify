package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 3v.29c2.892.86 5 3.539 5 6.71v7h1v2H4v-2h1v-7a7.003 7.003 0 0 1 5-6.71V3a2 2 0 1 1 4 0ZM7 17h10v-7a5 5 0 0 0-10 0v7Zm7 4v-1h-4v1a2 2 0 1 0 4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}