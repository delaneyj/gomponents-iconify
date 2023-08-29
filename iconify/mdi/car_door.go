package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarDoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14h-3v2h3v-2m3 7H3V11l8-8h10a1 1 0 0 1 1 1v17M11.83 5l-6 6H20V5h-8.17Z"/>`),
		g.Group(children),
	)
}