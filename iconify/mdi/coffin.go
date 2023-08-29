package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22L5 8l3-6h8l3 6l-3 14H8m3-16v2H9v2h2v5h2v-5h2V8h-2V6h-2Z"/>`),
		g.Group(children),
	)
}