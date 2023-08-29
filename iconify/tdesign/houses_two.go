package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HousesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2v2h9.566l4.2 7H21v11H3V11H1.233l4.2-7H7V2h2Zm-4 9v9h4v-5h6v5h4v-9H5Zm14.233-2l-1.8-3H6.566l-1.8 3h14.467ZM13 20v-3h-2v3h2Z"/>`),
		g.Group(children),
	)
}