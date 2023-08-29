package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothingStoreEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M2.5 1l-2 2v2h2v5h6V5h2V3l-2-2H7L5.5 4L4 1H2.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}