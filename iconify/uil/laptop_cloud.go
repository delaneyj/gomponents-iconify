package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 10H10a3 3 0 0 0 1.07-5.8a4 4 0 0 0-7.48 1A2.5 2.5 0 0 0 4.5 10Zm0-3a1 1 0 0 0 1-1a2 2 0 0 1 3.89-.64a1 1 0 0 0 .78.66A1 1 0 0 1 11 7a1 1 0 0 1-1 1H4.5a.5.5 0 0 1 0-1ZM21 16h-1V9a3 3 0 0 0-3-3h-1a1 1 0 0 0 0 2h1a1 1 0 0 1 1 1v7H6v-3a1 1 0 0 0-2 0v3H3a1 1 0 0 0-1 1v2a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-2a1 1 0 0 0-1-1Zm-1 3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16Z"/>`),
		g.Group(children),
	)
}