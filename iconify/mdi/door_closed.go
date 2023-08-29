package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 11h2v2h-2v-2m-4-8h7c1.11 0 2 .89 2 2v14h1v2H2v-2h8V5c0-1.11.89-2 2-2m0 2v14h7V5h-7Z"/>`),
		g.Group(children),
	)
}