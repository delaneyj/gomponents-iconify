package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeOneHundred(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.1 14.1L1 17l2 2l2.9-3.1L8 18v-6H2l2.1 2.1zM19 3l-2-2l-2.9 3.1L12 2v6h6l-2.1-2.1L19 3z"/>`),
		g.Group(children),
	)
}