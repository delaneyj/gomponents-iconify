package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.862a1 1 0 0 1 1.496-.868l5.492 3.138a1 1 0 0 1 0 1.736l-5.492 3.139A1 1 0 0 1 3 9.139V2.862ZM9.492 6L4 2.862v6.277L9.492 6Z"/>`),
		g.Group(children),
	)
}