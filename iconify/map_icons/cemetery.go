package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cemetery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25 1C15.062 1 7 8.909 7 18.664V49h36V18.664C43 8.909 34.944 1 25 1z"/>`),
		g.Group(children),
	)
}