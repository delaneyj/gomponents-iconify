package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 4.594L18.28 6.28l-9 9l-.686.72l.687.72l9 9L20 27.405V4.594zm-2 4.843v13.126L11.437 16L18 9.437z"/>`),
		g.Group(children),
	)
}