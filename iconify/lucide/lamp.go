package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 2h8l4 10H4L8 2Zm4 10v6m-4 4v-2c0-1.1.9-2 2-2h4a2 2 0 0 1 2 2v2H8Z"/>`),
		g.Group(children),
	)
}