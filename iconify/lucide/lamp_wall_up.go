package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampWallUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4h6l3 7H8l3-7Zm3 7v5a2 2 0 0 1-2 2H8m-4-3h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H4v-6Z"/>`),
		g.Group(children),
	)
}