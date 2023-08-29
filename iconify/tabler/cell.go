package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4L4 6v5l4 2l4-2V6zm4 7l4 2l4-2V6l-4-2l-4 2m-4 7v5l4 2l4-2v-5"/>`),
		g.Group(children),
	)
}