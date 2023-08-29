package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Import(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 4v24h20v-9h-2v7H8V6h16v7h2V4H6zm11.5 7l-4.313 4.28l-.687.72l.688.72L17.5 21l1.406-1.406L16.313 17H28v-2H16.312l2.594-2.594L17.5 11z"/>`),
		g.Group(children),
	)
}