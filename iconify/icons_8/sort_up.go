package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 8.594l-.72.687l-10 10L3.595 21h24.811l-1.687-1.72l-10-10l-.72-.686zm0 2.844L23.563 19H8.438L16 11.437z"/>`),
		g.Group(children),
	)
}