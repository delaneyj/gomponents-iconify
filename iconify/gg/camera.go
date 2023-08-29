package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4.5v2h8v1H3a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h12a3 3 0 0 0 2.99-2.751L24 17.5v-8l-6.01.751A3 3 0 0 0 15 7.5h-1v-2a1 1 0 0 0-1-1H4Zm14 7.766v2.468l4 .5v-3.468l-4 .5ZM16 10.5a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}