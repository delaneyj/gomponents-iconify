package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2v5.5l3 3l3-3V2H9m2 2h2v2h-2V4M2 9v6h5.5l3-3l-3-3H2m14.5 0l-3 3l3 3H22V9h-5.5M12 13.5l-3 3V22h6v-5.5l-3-3Z"/>`),
		g.Group(children),
	)
}