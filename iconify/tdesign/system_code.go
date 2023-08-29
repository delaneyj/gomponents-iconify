package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v17H1V1Zm2 2v13h18V3H3Zm7.406 3.844L8.28 9.5l2.125 2.656l-1.562 1.25L5.719 9.5l3.125-3.906l1.562 1.25Zm4.75-1.25L18.281 9.5l-3.125 3.906l-1.562-1.25L15.72 9.5l-2.125-2.656l1.562-1.25ZM3.222 21h17.556v2H3.222v-2Z"/>`),
		g.Group(children),
	)
}