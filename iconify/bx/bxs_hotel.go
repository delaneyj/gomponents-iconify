package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsHotel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="8" cy="11" r="3" fill="currentColor"/><path d="M18.205 7H12v8H4V6H2v14h2v-3h16v3h2v-4c0-.009-.005-.016-.005-.024H22V11c0-2.096-1.698-4-3.795-4z" fill="currentColor"/>`),
		g.Group(children),
	)
}