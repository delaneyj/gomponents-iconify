package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BankJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m12.28 3.376l-2.9 3.625h1.87a.75.75 0 0 1 0 1.5H8.5v1h2.75a.75.75 0 0 1 0 1.5H8.5v1.25a1 1 0 0 1-2 0v-1.25H3.75a.75.75 0 1 1 0-1.5H6.5v-1H3.75a.75.75 0 0 1 0-1.5h1.87l-2.9-3.625a1 1 0 0 1 1.557-1.254l.004.004L7.5 6.15l3.22-4.024a1 1 0 0 1 1.565 1.245Z"/>`),
		g.Group(children),
	)
}