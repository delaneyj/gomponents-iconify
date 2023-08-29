package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextformatBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3H5v18h9a5 5 0 0 0 2.436-9.367A5 5 0 0 0 13 3H6Zm7 8H7V5h6a3 3 0 1 1 0 6Zm-6 2h7a3 3 0 1 1 0 6H7v-6Z"/>`),
		g.Group(children),
	)
}