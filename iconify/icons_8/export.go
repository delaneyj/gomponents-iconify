package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 4v24h20v-8l-2 2v4H8V6h16v4l2 2V4H6zm16.406 7L21 12.406L23.563 15h-9.657v2h9.656L21 19.594L22.406 21l4.313-4.28l.686-.72l-.687-.72L22.405 11z"/>`),
		g.Group(children),
	)
}