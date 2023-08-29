package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 21a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h19ZM11 5H8.985v8h-1v6H12v-6h-1V5Zm7.015 14H22V5h-2.985v8h-1v6Zm-1-6h-1V5H14v8h-1v6h4.015v-6Zm-10.03 6v-6h-1V5H3v14h3.985Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}