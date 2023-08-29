package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BacktopRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm13.5 4h-11V6h11v2Zm-5.5.808l4.596 4.596l-1.414 1.414L13 12.636V18.5h-2v-5.864l-2.182 2.182l-1.414-1.414L12 8.808Z"/>`),
		g.Group(children),
	)
}