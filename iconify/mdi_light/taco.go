package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 18A2.5 2.5 0 0 1 1 15.5C1 11.36 4.81 8 9.5 8c.69 0 1.36.07 2 .21c.64-.14 1.31-.21 2-.21c4.69 0 8.5 3.36 8.5 7.5a2.5 2.5 0 0 1-2.5 2.5h-16M2 15.5A1.5 1.5 0 0 0 3.5 17A1.5 1.5 0 0 0 5 15.5c0-2.78 1.71-5.2 4.25-6.5C5.22 9.12 2 12 2 15.5M19.5 17a1.5 1.5 0 0 0 1.5-1.5c0-3.59-3.36-6.5-7.5-6.5C9.36 9 6 11.91 6 15.5c0 .56-.19 1.08-.5 1.5h14Z"/>`),
		g.Group(children),
	)
}