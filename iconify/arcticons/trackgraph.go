package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trackgraph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 20.22H27.78V5.5h-7.56v14.72H5.5v7.56h14.72V42.5h7.56V27.78H42.5v-7.56z"/>`),
		g.Group(children),
	)
}