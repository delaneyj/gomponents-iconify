package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12 2l4 4v12H4V2h8zM9 13l-2-2l2-2l-1-1l-3 3l3 3zm3 1l3-3l-3-3l-1 1l2 2l-2 2z"/>`),
		g.Group(children),
	)
}