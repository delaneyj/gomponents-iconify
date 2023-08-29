package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapeHexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6 15.235l6 3.333l6-3.333v-6.47l-6-3.333l-6 3.333v6.47ZM12 2L3 7v10l9 5l9-5V7l-9-5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}