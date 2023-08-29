package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 3.594l-.72.687l-8 8L5.595 14h20.811l-1.687-1.72l-8-8l-.72-.686zm0 2.844L21.563 12H10.438L16 6.437zM5.594 18l1.687 1.72l8 8l.72.686l.72-.687l8-8L26.405 18H5.594zm4.843 2h11.126L16 25.563L10.437 20z"/>`),
		g.Group(children),
	)
}