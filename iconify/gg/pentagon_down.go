package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 8L7 6v10l5 2.5l5-2.5V6l-5 2Zm3 .954l-3 1.2l-3-1.2v5.81l3 1.5l3-1.5v-5.81Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}