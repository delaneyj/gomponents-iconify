package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 3l-3.5 7H16l-6 12.03V14H7V3h9m-5.11 8l3.49-7H8v9h3v4.79L14.38 11h-3.49Z"/>`),
		g.Group(children),
	)
}